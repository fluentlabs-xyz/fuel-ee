package graphql_object

import (
	"github.com/graphql-go/graphql"
)

type CoinTypeType struct {
	SchemaFields UnionFields
}

//	pub enum CoinType {
//	   Coin(Coin),
//	   MessageCoin(MessageCoin),
//	   #[cynic(fallback)]
//	   Unknown,
//	}
type CoinTypeStruct struct {
}

func NewCoinTypeType(coinType *CoinType, messageCoinType *MessageCoinType) *CoinTypeType {
	config := graphql.UnionConfig{
		Name: "CoinType",
		Types: []*graphql.Object{
			coinType.SchemaFields.Object,
			messageCoinType.SchemaFields.Object,
		},
		ResolveType: func(p graphql.ResolveTypeParams) *graphql.Object {
			if _, ok := p.Value.(*CoinStruct); ok {
				return coinType.SchemaFields.Object
			}
			if _, ok := p.Value.(*MessageCoinStruct); ok {
				return messageCoinType.SchemaFields.Object
			}
			return nil
		},
	}
	entity := graphql.NewUnion(config)

	return &CoinTypeType{
		SchemaFields: UnionFields{
			Config: &config,
			Type:   entity,
		},
	}
}
