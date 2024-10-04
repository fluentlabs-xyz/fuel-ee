package graphql_object

import (
	"github.com/fluentlabs-xyz/fuel-ee/src/graphql_scalars"
	"github.com/graphql-go/graphql"
)

type SubmitType struct {
	SchemaFields SchemaFields
}

//		pub struct TransactionIdFragment {
//	   pub id: TransactionId,
//	}
type SubmitStruct struct {
	Id *graphql_scalars.Bytes32 `json:"id"`
}

func MakeSubmitType() (*SubmitType, error) {
	objectConfig := graphql.ObjectConfig{Name: "Submit", Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql_scalars.Bytes32Type,
			//Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			//	return 0, nil
			//},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	return &SubmitType{
		SchemaFields: SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
