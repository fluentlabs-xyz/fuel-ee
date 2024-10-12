package graphql_entrypoints

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fluentlabs-xyz/fuel-ee/src/config"
	"github.com/fluentlabs-xyz/fuel-ee/src/graphql_object"
	"github.com/fluentlabs-xyz/fuel-ee/src/graphql_scalars"
	"github.com/graphql-go/graphql"
)
import log "github.com/sirupsen/logrus"

type DryRunEntry struct {
	SchemaFields graphql_object.SchemaFields
}

type DryRunEntryStruct struct {
}

// const encodedTransactionsArgName = "encodedTransactions"
const encodedTransactionsArgName = "txs"
const utxoValidationArgName = "utxoValidation"
const gasPriceArgName = "gasPrice"

func MakeDryRunEntry(ethClient *ethclient.Client, dryRunTransactionStatusType *graphql_object.DryRunTransactionExecutionStatusType, config *config.Config) (*DryRunEntry, error) {
	objectConfig := graphql.ObjectConfig{Name: "DryRunEntry", Fields: graphql.Fields{
		"dryRun": &graphql.Field{
			Type: graphql.NewList(dryRunTransactionStatusType.SchemaFields.Object),
			Args: graphql.FieldConfigArgument{
				encodedTransactionsArgName: &graphql.ArgumentConfig{
					Type:         graphql.NewList(graphql_scalars.HexStringType),
					DefaultValue: []graphql_scalars.HexString{},
				},
				utxoValidationArgName: &graphql.ArgumentConfig{
					Type:         graphql.Boolean,
					DefaultValue: false,
				},
				gasPriceArgName: &graphql.ArgumentConfig{
					Type:         graphql_scalars.U64Type,
					DefaultValue: graphql_scalars.NewU64(0),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				encodedTransactions := p.Args[encodedTransactionsArgName]
				//utxoValidation := p.Args[utxoValidationArgName]
				//gasPrice := p.Args[gasPriceArgName]
				encodedTransactionsList, ok := encodedTransactions.([]interface{})
				if !ok {
					return nil, errors.New("encoded transactions must be a list")
				}
				results := make([]*graphql_object.DryRunTransactionExecutionStatusStruct, 0)
				for _, encodedTransaction := range encodedTransactionsList {
					transactionHexString, ok := encodedTransaction.(*graphql_scalars.HexString)
					if !ok {
						return nil, errors.New("each encoded transaction must be a hex string")
					}

					// send tx to reth node for emulation/estimation process (to get status, receipts, gas spent)
					from := common.HexToAddress(config.Blockchain.FuelRelayerAddress)
					to := common.HexToAddress(config.Blockchain.FuelContractAddress)
					callMsg := ethereum.CallMsg{
						From: from,
						To:   &to,
						Data: append(config.Blockchain.FvmDryRunSigBytes, transactionHexString.Value()...),
					}
					estimatedGas, err := ethClient.EstimateGas(context.Background(), callMsg)
					if err != nil {
						return nil, errors.New(fmt.Sprintf("DryRun: failed to estimate gas, error: %s", err))
					}
					callMsg.Gas = estimatedGas
					callRes, err := ethClient.CallContract(context.Background(), callMsg, nil)
					if err != nil {
						return nil, errors.New(fmt.Sprintf("DryRun: failed to call contract, error: %s", err))
					}
					log.Printf("DryRun: callRes: %s", callRes)
					results = append(results, &graphql_object.DryRunTransactionExecutionStatusStruct{
						Id: "0x0000000000000000000000000000000000000000000000000000000000000000",
						Status: &graphql_object.DryRunSuccessStatusStruct{
							TotalGas: int(estimatedGas),
							TotalFee: 0,
							ProgramState: &graphql_object.ProgramStateStruct{
								ReturnType: "RETURN",
								Data:       hex.EncodeToString(callRes),
							},
						},
						Receipts: []graphql_object.ReceiptStruct{},
					})
				}
				return results, nil
			},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{
		Query:    object,
		Mutation: object,
	}
	schema, err := graphql.NewSchema(schemaConfig)

	return &DryRunEntry{
		SchemaFields: graphql_object.SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
