package graphql_entrypoints

import (
	"github.com/fluentlabs-xyz/fuel-ee/src/graphql_object"
	"github.com/graphql-go/graphql"
)

type GetCoinsEntry struct {
	SchemaFields graphql_object.SchemaFields
}

//	{
//	 "data": {
//	   "coins": {
//	     "pageInfo": {
//	       "hasPreviousPage": false,
//	       "hasNextPage": false,
//	       "startCursor": "0x00000000000000000000000000000000000000000000000000000000000000010000",
//	       "endCursor": "0x00000000000000000000000000000000000000000000000000000000000000010000"
//	     },
//	     "edges": [
//	       {
//	         "node": {
//	           "type": "Coin",
//	           "utxoId": "0x00000000000000000000000000000000000000000000000000000000000000010000",
//	           "owner": "0x6b63804cfbf9856e68e5b6e7aef238dc8311ec55bec04df774003a2c96e0418e",
//	           "amount": "1152921504606846976",
//	           "assetId": "0xf8f8b6283d7fa5b672b530cbb84fcccb4ff8dc40f8176ef4544ddb1f1952ad07",
//	           "blockCreated": "0",
//	           "txCreatedIdx": "0"
//	         }
//	       }
//	     ]
//	   }
//	 }
//	}
type GetCoinsStruct struct {
}

func MakeGetCoinsEntry(pageInfoType *graphql_object.PageInfoType) (*GetCoinsEntry, error) {
	objectConfig := graphql.ObjectConfig{Name: "GetCoinsEntry", Fields: graphql.Fields{
		"pageInfo": &graphql.Field{
			Type: pageInfoType.SchemaFields.Object,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &graphql_object.PageInfoStruct{}, nil
			},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	return &GetCoinsEntry{
		SchemaFields: graphql_object.SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
