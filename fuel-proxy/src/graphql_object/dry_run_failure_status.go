package graphql_object

import (
	"github.com/graphql-go/graphql"
)

type DryRunFailureStatusType struct {
	SchemaFields SchemaFields
}

//	pub struct MakeDryRunFailureStatus {
//	   pub program_state: Option<ProgramState>,
//	   pub receipts: Vec<Receipt>,
//	   pub total_gas: U64,
//	   pub total_fee: U64,
//	}
type DryRunFailureStatusStruct struct {
	TotalGas     string              `json:"totalGas"`
	TotalFee     string              `json:"totalFee"`
	Reason       string              `json:"reason"`
	ProgramState *ProgramStateStruct `json:"programState"`
}

func MakeDryRunFailureStatus(programStateType *ProgramStateType) (*DryRunFailureStatusType, error) {
	objectConfig := graphql.ObjectConfig{Name: "DryRunFailureStatus", Fields: graphql.Fields{
		"totalGas": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
		"totalFee": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
		"reason": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "not set", nil // TODO
			},
		},
		"programState": &graphql.Field{
			Type: programStateType.SchemaFields.Object,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &ProgramStateStruct{
					ReturnType: "RETURN",             // TODO
					Data:       "0x0000000000000000", // TODO
				}, nil
			},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	return &DryRunFailureStatusType{
		SchemaFields: SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
