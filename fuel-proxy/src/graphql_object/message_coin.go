package graphql_object

import (
	"github.com/fluentlabs-xyz/fuel-ee/src/graphql_scalars"
	"github.com/graphql-go/graphql"
)

type MessageCoinType struct {
	SchemaFields SchemaFields
}

//	pub struct MessageCoin {
//	   pub amount: U64,
//	   pub nonce: Nonce,
//	   pub da_height: U64,
//	   pub sender: Address,
//	   pub recipient: Address,
//	}
type MessageCoinStruct struct {
	Amount    uint64                   `json:"amount"`
	Nonce     uint64                   `json:"nonce"`
	DaHeight  uint64                   `json:"daHeight"`
	Sender    *graphql_scalars.Bytes32 `json:"sender"`
	AssetId   *graphql_scalars.Bytes32 `json:"assetId"`
	Recipient *graphql_scalars.Bytes32 `json:"recipient"`
}

func NewMessageCoinType() (*MessageCoinType, error) {
	objectConfig := graphql.ObjectConfig{Name: "MessageCoin", Fields: graphql.Fields{
		"amount": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				coin, ok := p.Source.(*MessageCoinStruct)
				if ok {
					return coin.Amount, nil
				}
				return 0, nil
			},
		},
		"nonce": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				coin, ok := p.Source.(*MessageCoinStruct)
				if ok {
					return coin.Nonce, nil
				}
				return 0, nil
			},
		},
		"daHeight": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				coin, ok := p.Source.(*MessageCoinStruct)
				if ok {
					return coin.DaHeight, nil
				}
				return 0, nil
			},
		},
		"sender": &graphql.Field{
			Type: graphql_scalars.Bytes32Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				coin, ok := p.Source.(*MessageCoinStruct)
				if ok {
					return coin.Sender, nil
				}
				return graphql_scalars.NewBytes32TryFromStringOrPanic("0x0000000000000000000000000000000000000000000000000000000000000000"), nil
			},
		},
		"assetId": &graphql.Field{
			Type: graphql_scalars.Bytes32Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				coin, ok := p.Source.(*MessageCoinStruct)
				if ok {
					return coin.AssetId, nil
				}
				return graphql_scalars.NewBytes32TryFromStringOrPanic("0x0000000000000000000000000000000000000000000000000000000000000000"), nil
			},
		},
		"recipient": &graphql.Field{
			Type: graphql_scalars.Bytes32Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				coin, ok := p.Source.(*MessageCoinStruct)
				if ok {
					return coin.Recipient, nil
				}
				return graphql_scalars.NewBytes32TryFromStringOrPanic("0x0000000000000000000000000000000000000000000000000000000000000000"), nil
			},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	return &MessageCoinType{
		SchemaFields: SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
