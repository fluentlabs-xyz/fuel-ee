package graphql_entrypoints

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fluentlabs-xyz/fuel-ee/src/config"
	"github.com/fluentlabs-xyz/fuel-ee/src/graphql_object"
	"github.com/fluentlabs-xyz/fuel-ee/src/graphql_scalars"
	"github.com/fluentlabs-xyz/fuel-ee/src/types"
	"github.com/graphql-go/graphql"
	log "github.com/sirupsen/logrus"
	"time"
)

type SubmitEntry struct {
	SchemaFields graphql_object.SchemaFields
}

type SubmitEntryStruct struct {
}

// const encodedTransactionsArgName = "encodedTransactions"
const encodedTransactionArgName = "tx"

func MakeSubmitEntry(ethClient *ethclient.Client, submitType *graphql_object.SubmitType, config *config.Config) (*SubmitEntry, error) {
	objectConfig := graphql.ObjectConfig{Name: "SubmitEntry", Fields: graphql.Fields{
		"submit": &graphql.Field{
			Type: submitType.SchemaFields.Object,
			Args: graphql.FieldConfigArgument{
				encodedTransactionArgName: &graphql.ArgumentConfig{
					Type:         graphql_scalars.HexStringType,
					DefaultValue: []graphql_scalars.HexString{},
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				encodedTransaction := p.Args[encodedTransactionArgName]
				transactionHexString, ok := encodedTransaction.(*graphql_scalars.HexString)
				if !ok {
					return nil, errors.New("each encoded transaction must be a hex string")
				}
				log.Printf("Submit: transactionHexString: %s", transactionHexString)

				// send tx to reth node for emulation/estimation process (to get status, receipts, gas spent)
				data := transactionHexString.Value()
				from := common.HexToAddress(types.FuelRelayerAccountAddress)
				to := common.HexToAddress(types.EthFuelVMPrecompileAddress)
				chainId, err := ethClient.ChainID(context.Background())
				if err != nil {
					return nil, errors.New(fmt.Sprintf("Submit: failed to fetch chain id, error: %s", err))
				}
				nonce, err := ethClient.PendingNonceAt(context.Background(), from)
				if err != nil {
					return nil, errors.New(fmt.Sprintf("Submit: failed to fetch pending nonce, error: %s", err))
				}
				gasPrice, err := ethClient.SuggestGasPrice(context.Background())
				if err != nil {
					return nil, errors.New(fmt.Sprintf("Submit: failed to estimate gas price, error: %s", err))
				}
				tx := ethereumTypes.NewTx(&ethereumTypes.AccessListTx{
					Gas:      300_000_000,
					GasPrice: gasPrice,
					To:       &to,
					Nonce:    nonce,
					Data:     append(config.Blockchain.FvmExecSigBytes, data...),
				})
				log.Printf("Submit: tx.Hash(): %s tx.Hash().String(): %s", tx.Hash(), tx.Hash().String())
				privateKey, err := crypto.HexToECDSA(config.Relayer.PrivateKey)
				if err != nil {
					return nil, errors.New(fmt.Sprintf("Submit: failed to create private key, error: %s", err))
				}
				signedTx, err := ethereumTypes.SignTx(tx, ethereumTypes.NewEIP2930Signer(chainId), privateKey)
				if err != nil {
					return nil, errors.New(fmt.Sprintf("Submit: failed to sign tx, error: %s", err))
				}
				log.Printf("Submit: len(signedTx.Data()) %d", len(signedTx.Data()))
				err = ethClient.SendTransaction(context.Background(), signedTx)
				if err != nil {
					return nil, errors.New(fmt.Sprintf("Submit: failed to send tx, error: %s", err))
				}
				log.Printf("Submit: signedTx.Hash().Hex(): %s signedTx.To(): %s", signedTx.Hash().Hex(), signedTx.To())
				isPending := true
				for isPending {
					tx, isPending, err = ethClient.TransactionByHash(context.Background(), signedTx.Hash())
					if err != nil {
						log.Printf("Submit: TransactionByHash, error: %s", err)
						//break
					} else {
						log.Printf(
							"Submit: tx: isPending:%v Hash:%s Nonce:%d ChainId:%s To:%s",
							isPending,
							tx.Hash(),
							tx.Nonce(),
							tx.ChainId(),
							tx.To(),
						)
						if !isPending {
							break
						}
					}
					time.Sleep(5 * time.Second)
				}
				return graphql_object.SubmitStruct{
					Id: graphql_scalars.NewBytes32TryFromStringOrPanic(signedTx.Hash().String()),
				}, nil
			},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{
		Query:    object,
		Mutation: object,
	}
	schema, err := graphql.NewSchema(schemaConfig)

	return &SubmitEntry{
		SchemaFields: graphql_object.SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
