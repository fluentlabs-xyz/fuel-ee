package graphql_entrypoints

import (
	"github.com/fluentlabs-xyz/fuel-ee/graphql_schemas"
	"github.com/fluentlabs-xyz/fuel-ee/graphql_types"
	"github.com/graphql-go/graphql"
)

type EstimateGasPriceType struct {
	SchemaFields graphql_schemas.SchemaFields
}

type EstimateGasPriceStruct struct {
}

const BlockHorizon = "blockHorizon"

func EstimateGasPrice(gasPriceType *graphql_schemas.GasPriceType) (*EstimateGasPriceType, error) {
	objectConfig := graphql.ObjectConfig{Name: "EstimateGasPrice", Fields: graphql.Fields{
		"estimateGasPrice": &graphql.Field{
			Type: gasPriceType.SchemaFields.Object,
			Args: graphql.FieldConfigArgument{
				BlockHorizon: &graphql.ArgumentConfig{
					Type:         graphql_types.U32Type,
					DefaultValue: graphql_types.NewU32(0),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				//blockHorizon := p.Args[BlockHorizon]
				return graphql_schemas.GasPriceStruct{
					GasPrice: 0,
				}, nil
			},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	return &EstimateGasPriceType{
		SchemaFields: graphql_schemas.SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
