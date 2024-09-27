package graphql_object

import "github.com/graphql-go/graphql"

type ReceiptTypeType struct {
	SchemaFields EnumFields
}

//		pub enum ReceiptType {
//	   Call,
//	   Return,
//	   ReturnData,
//	   Panic,
//	   Revert,
//	   Log,
//	   LogData,
//	   Transfer,
//	   TransferOut,
//	   ScriptResult,
//	   MessageOut,
//	   Mint,
//	   Burn,
//	}
type ReceiptTypeStruct struct {
}

func MakeReceiptType() *ReceiptTypeType {
	enumConfig := graphql.EnumConfig{
		Name: "ReceiptType",
		Values: graphql.EnumValueConfigMap{
			"Call": &graphql.EnumValueConfig{
				Value: 1,
			},
			"Return": &graphql.EnumValueConfig{
				Value: 2,
			},
			"ReturnData": &graphql.EnumValueConfig{
				Value: 3,
			},
			"Panic": &graphql.EnumValueConfig{
				Value: 4,
			},
			"Revert": &graphql.EnumValueConfig{
				Value: 5,
			},
			"Log": &graphql.EnumValueConfig{
				Value: 6,
			},
			"LogData": &graphql.EnumValueConfig{
				Value: 7,
			},
			"Transfer": &graphql.EnumValueConfig{
				Value: 8,
			},
			"TransferOut": &graphql.EnumValueConfig{
				Value: 9,
			},
			"ScriptResult": &graphql.EnumValueConfig{
				Value: 10,
			},
			"MessageOut": &graphql.EnumValueConfig{
				Value: 11,
			},
			"Mint": &graphql.EnumValueConfig{
				Value: 12,
			},
			"Burn": &graphql.EnumValueConfig{
				Value: 13,
			},
		},
	}
	enum := graphql.NewEnum(enumConfig)

	return &ReceiptTypeType{
		SchemaFields: EnumFields{
			Config: &enumConfig,
			Type:   enum,
		},
	}
}
