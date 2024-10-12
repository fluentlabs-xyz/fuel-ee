package graphql_object

import (
	"github.com/graphql-go/graphql"
)

type DryRunTransactionExecutionStatusType struct {
	SchemaFields SchemaFields
}

//	pub struct NewDryRunTransactionExecutionStatusType {
//	   pub id: TransactionId,
//	   pub status: MakeDryRunTransactionStatus,
//	}
type DryRunTransactionExecutionStatusStruct struct {
	Id       string          `json:"id"`
	Status   interface{}     `json:"status"`
	Receipts []ReceiptStruct `json:"receipts"`
}

func NewDryRunTransactionExecutionStatusType(dryRunTransactionStatusType *DryRunTransactionStatusType, receiptType *ReceiptType) (*DryRunTransactionExecutionStatusType, error) {
	objectConfig := graphql.ObjectConfig{Name: "DryRunTransactionExecutionStatus", Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if status, ok := p.Source.(*DryRunTransactionExecutionStatusStruct); ok {
					return status.Id, nil
				}
				return "0x0000000000000000000000000000000000000000000000000000000000000000", nil
			},
		},
		"status": &graphql.Field{
			Type: dryRunTransactionStatusType.SchemaFields.Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if status, ok := p.Source.(*DryRunTransactionExecutionStatusStruct); ok {
					return status.Status, nil
				}
				return &DryRunSuccessStatusStruct{
					TotalGas: 0,
					TotalFee: 0,
					ProgramState: &ProgramStateStruct{
						ReturnType: "RETURN",
						Data:       "0x0000000000000000",
					},
				}, nil
			},
		},
		"receipts": &graphql.Field{
			Type: graphql.NewList(receiptType.SchemaFields.Object),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if status, ok := p.Source.(*DryRunTransactionExecutionStatusStruct); ok {
					return status.Receipts, nil
				}
				return []ReceiptStruct{}, nil
			},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	return &DryRunTransactionExecutionStatusType{
		SchemaFields: SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
