package graphql_object

import "github.com/graphql-go/graphql"

type DryRunType struct {
	SchemaFields SchemaFields
}

//		pub struct DryRun {
//	   #[arguments(txs: $txs, utxoValidation: $utxo_validation, gasPrice: $gas_price)]
//	   pub dry_run: Vec<NewDryRunTransactionExecutionStatusType>,
//	}
type DryRunStruct struct {
	DryRun []DryRunTransactionExecutionStatusStruct `json:"dryRun"`
}

func DryRun(dryRunTransactionExecutionStatusType *DryRunTransactionExecutionStatusType) (*DryRunType, error) {
	objectConfig := graphql.ObjectConfig{Name: "DryRun", Fields: graphql.Fields{
		"dryRun": &graphql.Field{
			Type: graphql.NewList(dryRunTransactionExecutionStatusType.SchemaFields.Object),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	return &DryRunType{
		SchemaFields: SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
