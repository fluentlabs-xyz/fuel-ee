package graphql_entrypoints

import (
	"github.com/fluentlabs-xyz/fuel-ee/graphql_object"
	"github.com/fluentlabs-xyz/fuel-ee/graphql_scalars"
	"github.com/graphql-go/graphql"
)

type EstimateGasPriceEntry struct {
	SchemaFields graphql_object.SchemaFields
}

type EstimateGasPriceStruct struct {
}

const blockHorizon = "blockHorizon"

func MakeEstimateGasPriceEntry(gasPriceType *graphql_object.GasPriceType) (*EstimateGasPriceEntry, error) {
	objectConfig := graphql.ObjectConfig{Name: "EstimateGasPriceEntry", Fields: graphql.Fields{
		"estimateGasPrice": &graphql.Field{
			Type: gasPriceType.SchemaFields.Object,
			Args: graphql.FieldConfigArgument{
				blockHorizon: &graphql.ArgumentConfig{
					Type:         graphql_scalars.U32Type,
					DefaultValue: graphql_scalars.NewU32(0),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				//blockHorizon := p.Args[blockHorizon]
				return graphql_object.GasPriceStruct{
					GasPrice: 0,
				}, nil
			},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	return &EstimateGasPriceEntry{
		SchemaFields: graphql_object.SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
