package graphql_entrypoints

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	types2 "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fluentlabs-xyz/fuel-ee/graphql_object"
	"github.com/fluentlabs-xyz/fuel-ee/graphql_scalars"
	"github.com/fluentlabs-xyz/fuel-ee/types"
	"github.com/graphql-go/graphql"
	"log"
)

type SubmitEntry struct {
	SchemaFields graphql_object.SchemaFields
}

type SubmitEntryStruct struct {
}

// const encodedTransactionsArgName = "encodedTransactions"
const encodedTransactionArgName = "tx"

func MakeSubmitEntry(ethClient *ethclient.Client, submitType *graphql_object.SubmitType) (*SubmitEntry, error) {
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
				log.Printf("transactionHexString: %s", transactionHexString)

				fvmSubmitSigBytes := make([]byte, 4)
				binary.BigEndian.PutUint32(fvmSubmitSigBytes, types.FvmExecSig)
				// send tx to reth node for emulation/estimation process (to get status, receipts, gas spent)
				//data := append(fvmSubmitSigBytes, transactionHexString.Value()...)
				data := transactionHexString.Value()
				from := common.HexToAddress(types.FuelRelayerAccountAddress)
				to := common.HexToAddress(types.EthFuelVMPrecompileAddress)
				callMsg := ethereum.CallMsg{
					From:          from,
					To:            &to,
					Gas:           0,
					Data:          data,
					AccessList:    nil,
					BlobGasFeeCap: nil,
					BlobHashes:    nil,
				}
				estimatedGas, err := ethClient.EstimateGas(context.Background(), callMsg)
				if err != nil {
					return nil, errors.New(fmt.Sprintf("Submit: failed to estimate gas, error: %s", err))
				}
				log.Printf("estimatedGas: %d", estimatedGas)
				tx := types2.NewTx(&types2.DynamicFeeTx{
					Gas:  estimatedGas,
					To:   &to,
					Data: data,
				})
				err = ethClient.SendTransaction(context.Background(), tx)
				if err != nil {
					return nil, errors.New(fmt.Sprintf("Submit: failed to call contract, error: %s", err))
				}
				return graphql_object.SubmitStruct{
					// TODO provide id
					Id: graphql_scalars.NewBytes32TryFromStringOrPanic("0x1231231230000000000000000000000000000000000000000000000000123123"),
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
