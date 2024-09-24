package graphql_schemas

import (
	"github.com/graphql-go/graphql"
)

type DependentCostType struct {
	SchemaFields UnionFields
}

//	pub enum DependentCost {
//	   LightOperation(LightOperation),
//	   HeavyOperation(HeavyOperation),
//	   #[cynic(fallback)]
//	   Unknown,
//	}
type DependentCostStruct struct {
}

func DependentCost(lightOperationType *LightOperationType, heavyOperationType *HeavyOperationType) *DependentCostType {
	config := graphql.UnionConfig{
		Name: "DependentCost",
		Types: []*graphql.Object{
			lightOperationType.SchemaFields.Object,
			heavyOperationType.SchemaFields.Object,
		},
		ResolveType: func(p graphql.ResolveTypeParams) *graphql.Object {
			if _, ok := p.Value.(*LightOperationStruct); ok {
				return lightOperationType.SchemaFields.Object
			}
			if _, ok := p.Value.(*HeavyOperationStruct); ok {
				return heavyOperationType.SchemaFields.Object
			}
			return nil
		},
	}
	entity := graphql.NewUnion(config)

	return &DependentCostType{
		SchemaFields: UnionFields{
			Config: &config,
			Entity: entity,
		},
	}
}
