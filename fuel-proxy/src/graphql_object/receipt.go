package graphql_object

import (
	"github.com/fluentlabs-xyz/fuel-ee/src/graphql_scalars"
	"github.com/graphql-go/graphql"
)

type ReceiptType struct {
	SchemaFields SchemaFields
}

//		pub struct Receipt {
//	   pub param1: Option<U64>,
//	   pub param2: Option<U64>,
//	   pub amount: Option<U64>,
//	   pub assetId: Option<AssetId>,
//	   pub gas: Option<U64>,
//	   pub digest: Option<Bytes32>,
//	   pub id: Option<ContractId>,
//	   pub is: Option<U64>,
//	   pub pc: Option<U64>,
//	   pub ptr: Option<U64>,
//	   pub ra: Option<U64>,
//	   pub rb: Option<U64>,
//	   pub rc: Option<U64>,
//	   pub rd: Option<U64>,
//	   pub reason: Option<U64>,
//	   pub receipt_type: ReceiptType,
//	   pub to: Option<ContractId>,
//	   pub to_address: Option<Address>,
//	   pub val: Option<U64>,
//	   pub len: Option<U64>,
//	   pub result: Option<U64>,
//	   pub gas_used: Option<U64>,
//	   pub data: Option<HexString>,
//	   pub sender: Option<Address>,
//	   pub recipient: Option<Address>,
//	   pub nonce: Option<Nonce>,
//	   pub contract_id: Option<ContractId>,
//	   pub sub_id: Option<Bytes32>,
//	}
type ReceiptStruct struct {
	// id
	//  pc
	//  is
	//  to
	//  toAddress
	//  amount
	//  assetId
	//  gas
	//  param1
	//  param2
	//  val
	//  ptr
	//  digest
	//  reason
	//  ra
	//  rb
	//  rc
	//  rd
	//  len
	//  receiptType
	//  result
	//  gasUsed
	//  data
	//  sender
	//  recipient
	//  nonce
	//  contractId
	//  subId
}

func Receipt(receiptTypeType *ReceiptTypeType) (*ReceiptType, error) {
	objectConfig := graphql.ObjectConfig{Name: "Receipt", Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql_scalars.Bytes32Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
		"pc": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
		"is": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
		//  to
		"to": &graphql.Field{
			Type: graphql_scalars.Bytes32Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
		//  toAddress
		"toAddress": &graphql.Field{
			Type: graphql_scalars.Bytes32Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
		//  amount
		"amount": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
		//  assetId
		"assetId": &graphql.Field{
			Type: graphql_scalars.Bytes32Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
		//  gas
		"gas": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
		//  param1
		"param1": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
		//  param2
		"param2": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
		//  val
		"val": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
		//  ptr
		"ptr": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
		//  digest
		"digest": &graphql.Field{
			Type: graphql_scalars.Bytes32Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
		//  reason
		"reason": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
		//  ra
		"ra": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
		//  rb
		"rb": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
		//  rc
		"rc": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
		//  rd
		"rd": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
		//  len
		"len": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
		//  receiptType
		"receiptType": &graphql.Field{
			Type: receiptTypeType.SchemaFields.Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
		//  result
		"result": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
		//  gasUsed
		"gasUsed": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
		//  data
		"data": &graphql.Field{
			Type: graphql_scalars.HexStringType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "", nil
			},
		},
		//  sender
		"sender": &graphql.Field{
			Type: graphql_scalars.Bytes32Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
		//  recipient
		"recipient": &graphql.Field{
			Type: graphql_scalars.Bytes32Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
		//  nonce
		"nonce": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
		//  contractId
		"contractId": &graphql.Field{
			Type: graphql_scalars.Bytes32Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
		//  subId
		"subId": &graphql.Field{
			Type: graphql_scalars.Bytes32Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	return &ReceiptType{
		SchemaFields: SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
