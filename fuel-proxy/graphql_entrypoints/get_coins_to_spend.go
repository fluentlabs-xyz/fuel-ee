package graphql_entrypoints

import (
	"github.com/fluentlabs-xyz/fuel-ee/graphql_input_objects"
	"github.com/fluentlabs-xyz/fuel-ee/graphql_object"
	"github.com/fluentlabs-xyz/fuel-ee/graphql_scalars"
	"github.com/graphql-go/graphql"
	"log"
)

type GetCoinsToSpendEntry struct {
	SchemaFields graphql_object.SchemaFields
}

//	{
//	 "data": {
//	   "coinsToSpend": [
//	     [
//	       {
//	         "type": "Coin",
//	         "utxoId": "0xce1e6751f5a4bbb53e3f63e9f8bfcf52281429a862d25e862e45c0cafcbf8daa0001",
//	         "owner": "0x6b63804cfbf9856e68e5b6e7aef238dc8311ec55bec04df774003a2c96e0418e",
//	         "amount": "1152921504606846974",
//	         "assetId": "0xf8f8b6283d7fa5b672b530cbb84fcccb4ff8dc40f8176ef4544ddb1f1952ad07",
//	         "blockCreated": "2",
//	         "txCreatedIdx": "0"
//	       }
//	     ]
//	   ]
//	 }
//	}
type GetCoinsToSpendStruct struct {
}

const ownerArgName = "owner"
const queryPerAssetArgName = "queryPerAsset"
const excludedIdsArgName = "excludedIds"

func MakeGetCoinsToSpendEntry(
	coinTypeType *graphql_object.CoinTypeType,
) (*GetCoinsToSpendEntry, error) {
	objectConfig := graphql.ObjectConfig{Name: "GetCoinsToSpendEntry", Fields: graphql.Fields{
		"coinsToSpend": &graphql.Field{
			Type: graphql.NewList(graphql.NewList(coinTypeType.SchemaFields.Type)),
			Args: graphql.FieldConfigArgument{
				ownerArgName: &graphql.ArgumentConfig{
					Type: graphql_scalars.AddressType,
					//DefaultValue: []graphql_scalars.HexString{},
				},
				queryPerAssetArgName: &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql_object.SpendQueryElementInput),
					//DefaultValue: []graphql_scalars.HexString{},
				},
				excludedIdsArgName: &graphql.ArgumentConfig{
					Type: graphql_input_objects.ExcludeInput,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				owner := p.Args[ownerArgName]
				queryPerAsset := p.Args[queryPerAssetArgName]
				excludedIds := p.Args[excludedIdsArgName]
				log.Printf("owner: %s", owner)
				log.Printf("queryPerAsset: %s", queryPerAsset)
				log.Printf("excludedIds: %s", excludedIds)
				// {
				//  "data": {
				//    "coinsToSpend": [
				//      [
				//        {
				//          "type": "Coin",
				//          "utxoId": "0xa9d5261a68ec08433015f7747d88d0541ced59213224fb96e5ba33e303314afb0001",
				//          "owner": "0x6b63804cfbf9856e68e5b6e7aef238dc8311ec55bec04df774003a2c96e0418e",
				//          "amount": "1152921504606846975",
				//          "assetId": "0xf8f8b6283d7fa5b672b530cbb84fcccb4ff8dc40f8176ef4544ddb1f1952ad07",
				//          "blockCreated": "1",
				//          "txCreatedIdx": "0"
				//        }
				//      ]
				//    ]
				//  }
				// }
				// TODO
				coin := graphql_object.CoinStruct{
					UtxoId:       *graphql_scalars.NewBytes34TryFromStringOrPanic("0xa9d5261a68ec08433015f7747d88d0541ced59213224fb96e5ba33e303314afb0001"),
					Owner:        *graphql_scalars.NewBytes32TryFromStringOrPanic("0x6b63804cfbf9856e68e5b6e7aef238dc8311ec55bec04df774003a2c96e0418e"),
					AssetId:      *graphql_scalars.NewBytes32TryFromStringOrPanic("0xf8f8b6283d7fa5b672b530cbb84fcccb4ff8dc40f8176ef4544ddb1f1952ad07"),
					Amount:       1152921504606846975,
					BlockCreated: 1,
					TxCreatedIdx: 0,
				}
				res := [][]*graphql_object.CoinStruct{{&coin}}
				return res, nil
			},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	return &GetCoinsToSpendEntry{
		SchemaFields: graphql_object.SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
