package graphql_entrypoints

import (
	"github.com/fluentlabs-xyz/fuel-ee/graphql_schemas"
	"github.com/graphql-go/graphql"
)

type GetCoinsType struct {
	SchemaFields graphql_schemas.SchemaFields
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

func GetCoins(pageInfoType *graphql_schemas.PageInfoType) (*GetCoinsType, error) {
	objectConfig := graphql.ObjectConfig{Name: "GetNodeInfo", Fields: graphql.Fields{
		"pageInfo": &graphql.Field{
			Type: pageInfoType.SchemaFields.Object,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &graphql_schemas.PageInfoStruct{}, nil
			},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	return &GetCoinsType{
		SchemaFields: graphql_schemas.SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
