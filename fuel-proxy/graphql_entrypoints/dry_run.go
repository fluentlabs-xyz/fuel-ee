package graphql_entrypoints

import (
	"github.com/fluentlabs-xyz/fuel-ee/graphql_object"
	"github.com/fluentlabs-xyz/fuel-ee/graphql_scalars"
	"github.com/graphql-go/graphql"
	"log"
)

type DryRunEntry struct {
	SchemaFields graphql_object.SchemaFields
}

type DryRunEntryStruct struct {
}

// const encodedTransactionsArgName = "encodedTransactions"
const encodedTransactionsArgName = "txs"
const utxoValidationArgName = "utxoValidation"
const gasPriceArgName = "gasPrice"

func MakeDryRunEntry(dryRunTransactionStatusType *graphql_object.DryRunTransactionExecutionStatusType) (*DryRunEntry, error) {
	objectConfig := graphql.ObjectConfig{Name: "DryRun", Fields: graphql.Fields{
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
				log.Printf("encodedTransactions: %s", encodedTransactions)
				return []graphql_object.DryRunTransactionExecutionStatusStruct{
					{
						Id:       "0xb4f5b359704eda15f8ec6c15004b6816b9df4f730baaa50d0a2fb34a99108bee",
						Status:   &graphql_object.DryRunTransactionStatusStruct{},
						Receipts: []graphql_object.ReceiptStruct{},
					},
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

	return &DryRunEntry{
		SchemaFields: graphql_object.SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
