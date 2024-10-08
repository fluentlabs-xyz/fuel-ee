package graphql_object

import (
	"github.com/graphql-go/graphql"
)

type GetCoinsType struct {
	SchemaFields SchemaFields
}

type GetCoinsStruct struct {
	PageInfo *PageInfoStruct `json:"pageInfo"`
	Edges    []*NodeStruct   `json:"edges"`
}

func NewGetCoinsType(pageInfoType *PageInfoType, coinType *CoinType) (*GetCoinsType, error) {
	nodeType, err := NewNodeType(coinType.SchemaFields.Object)
	if err != nil {
		return nil, err
	}
	objectConfig := graphql.ObjectConfig{Name: "GetCoins", Fields: graphql.Fields{
		"pageInfo": &graphql.Field{
			Type: pageInfoType.SchemaFields.Object,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				entity1, ok := p.Source.(*PageInfoStruct)
				if ok {
					return entity1, nil
				}
				entity2, ok := p.Source.(*GetCoinsStruct)
				if ok {
					return entity2.PageInfo, nil
				}
				return &PageInfoStruct{}, nil
			},
		},
		"edges": &graphql.Field{
			Type: graphql.NewList(nodeType.SchemaFields.Object),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				entity1, ok := p.Source.([]*NodeStruct)
				if ok {
					return entity1, nil
				}
				entity2, ok := p.Source.(*GetCoinsStruct)
				if ok {
					return entity2.Edges, nil
				}
				return []*NodeStruct{{
					Node: &CoinStruct{},
				}}, nil
			},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	return &GetCoinsType{
		SchemaFields: SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
