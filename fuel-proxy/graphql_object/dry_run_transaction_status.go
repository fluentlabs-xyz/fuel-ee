package graphql_object

import (
	"github.com/graphql-go/graphql"
)

type DryRunTransactionStatusType struct {
	SchemaFields UnionFields
}

//		pub enum MakeDryRunTransactionStatus {
//	   SuccessStatus(MakeDryRunSuccessStatus),
//	   FailureStatus(MakeDryRunFailureStatus),
//	   #[cynic(fallback)]
//	   Unknown,
//	}
type DryRunTransactionStatusStruct struct {
}

func MakeDryRunTransactionStatus(dryRunSuccessStatusType *DryRunSuccessStatusType, dryRunFailureStatusType *DryRunFailureStatusType) *DryRunTransactionStatusType {
	config := graphql.UnionConfig{
		Name: "DryRunTransactionStatus",
		Types: []*graphql.Object{
			dryRunSuccessStatusType.SchemaFields.Object,
			dryRunFailureStatusType.SchemaFields.Object,
		},
		ResolveType: func(p graphql.ResolveTypeParams) *graphql.Object {
			if _, ok := p.Value.(*DryRunSuccessStatusStruct); ok {
				return dryRunSuccessStatusType.SchemaFields.Object
			}
			if _, ok := p.Value.(*DryRunFailureStatusStruct); ok {
				return dryRunFailureStatusType.SchemaFields.Object
			}
			return nil
		},
	}
	entity := graphql.NewUnion(config)

	return &DryRunTransactionStatusType{
		SchemaFields: UnionFields{
			Config: &config,
			Type:   entity,
		},
	}
}
