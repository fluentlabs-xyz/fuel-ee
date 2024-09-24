package graphql_schemas

import "github.com/graphql-go/graphql"

type TxParametersType struct {
	SchemaFields SchemaFields
}

//	pub struct TxParameters {
//	   pub version: TxParametersVersion,
//	   pub max_inputs: U16,
//	   pub max_outputs: U16,
//	   pub max_witnesses: U32,
//	   pub max_gas_per_tx: U64,
//	   pub max_size: U64,
//	   pub max_bytecode_subsections: U16,
//	}
type TxParametersStruct struct {
	MaxBytecodeSubsections string `json:"maxBytecodeSubsections"`
	MaxGasPerTx            string `json:"maxGasPerTx"`
	MaxInputs              string `json:"maxInputs"`
	MaxOutputs             string `json:"maxOutputs"`
	MaxSize                string `json:"maxSize"`
	MaxWitnesses           string `json:"maxWitnesses"`
}

func TxParameters(txParametersVersionType *TxParametersVersionType) (*TxParametersType, error) {
	objectConfig := graphql.ObjectConfig{Name: "TxParameters", Fields: graphql.Fields{
		"version": &graphql.Field{
			Type: txParametersVersionType.SchemaFields.Entity,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 1, nil
			},
		},
		"maxInputs": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 255, nil
			},
		},
		"maxOutputs": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 255, nil
			},
		},
		"maxWitnesses": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 255, nil
			},
		},
		"maxGasPerTx": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 30000000, nil
			},
		},
		"maxSize": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 110 * 1024, nil
			},
		},
		"maxBytecodeSubsections": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 256, nil
			},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	return &TxParametersType{
		SchemaFields: SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
