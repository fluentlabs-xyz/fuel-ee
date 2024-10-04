package graphql_entrypoints

import (
	"context"
	"errors"
	"fmt"
	"github.com/fluentlabs-xyz/fuel-ee/src/graphql_input_objects"
	"github.com/fluentlabs-xyz/fuel-ee/src/graphql_object"
	"github.com/fluentlabs-xyz/fuel-ee/src/graphql_scalars"
	"github.com/fluentlabs-xyz/fuel-ee/src/services/utxoService"
	"github.com/graphql-go/graphql"
	log "github.com/sirupsen/logrus"
	"slices"
	"strings"
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
	utxoService *utxoService.Service,
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
				ownerArg := p.Args[ownerArgName]
				queryPerAssetArg := p.Args[queryPerAssetArgName]
				excludedIdsArg := p.Args[excludedIdsArgName]
				log.Printf("ownerArg: %s", ownerArg)
				log.Printf("queryPerAssetArg: %s", queryPerAssetArg)
				log.Printf("excludedIdsArg: %s", excludedIdsArg)
				owner, ok := ownerArg.(*graphql_scalars.Address)
				if !ok {
					errText := "param [ownerArg] must be of *graphql_scalars.Address type"
					log.Printf(errText)
					return nil, errors.New(errText)
				}
				excludedIdsArgInterfaces, ok := excludedIdsArg.(map[string]interface{})
				if !ok {
					errText := "param [excludedIdsArg] must be a map of interfaces"
					log.Printf(errText)
					return nil, errors.New(errText)
				}
				var excludedIds []*graphql_scalars.Bytes34
				excludedIdsInterfaces, ok := excludedIdsArgInterfaces["utxos"].([]interface{})
				if !ok {
					errText := "param [excludedIdsArgInterfaces] must have 'utxos' key which must be of type []interface{}"
					log.Printf(errText)
					return nil, errors.New(errText)
				}
				if len(excludedIdsInterfaces) > 0 {
					excludedIds, ok = excludedIdsArgInterfaces["utxos"].([]*graphql_scalars.Bytes34)
					if !ok {
						errText := "param [excludedIdsArgInterfaces] must have 'utxosInterfaces' key which values are of []*graphql_scalars.Bytes34 type"
						log.Printf(errText)
						return nil, errors.New(errText)
					}
				}
				excludedIdsStrings := make([]string, 0, len(excludedIds))
				if len(excludedIds) > 0 {
					for _, excludedId := range excludedIds {
						excludedIdsStrings = append(excludedIdsStrings, excludedId.String())
					}
				}
				// {
				//  "data": {
				//    "coinsToSpend": [
				//      [
				//        {
				//          "type": "Coin",
				//          "utxoId": "0xa9d5261a68ec08433015f7747d88d0541ced59213224fb96e5ba33e303314afb0001",
				//          "ownerArg": "0x6b63804cfbf9856e68e5b6e7aef238dc8311ec55bec04df774003a2c96e0418e",
				//          "amount": "1152921504606846975",
				//          "assetId": "0xf8f8b6283d7fa5b672b530cbb84fcccb4ff8dc40f8176ef4544ddb1f1952ad07",
				//          "blockCreated": "1",
				//          "txCreatedIdx": "0"
				//        }
				//      ]
				//    ]
				//  }
				// }
				// TODO query from some DB
				utxos, err := utxoService.Repo().FindAllByParams(context.Background(), strings.ToLower(owner.String()), "*", "*", false)
				if err != nil {
					errText := "param [excludedIdsArgInterfaces] must have 'utxosInterfaces' key which values are of []interface{} type"
					log.Printf(errText)
					return nil, errors.New(errText)
				}
				res := make([][]*graphql_object.CoinStruct, 0)
				for _, utxo := range utxos {
					if slices.Contains(excludedIdsStrings, utxo.TxId) {
						continue
					}
					utxoId, err := utxo.UtxoId()
					if err != nil {
						errText := fmt.Sprintf("failed to get utxo id, error: %s", err)
						log.Printf(errText)
						return nil, errors.New(errText)
					}
					coin := graphql_object.CoinStruct{
						UtxoId:       utxoId,
						Owner:        graphql_scalars.NewBytes32TryFromStringOrPanic(utxo.Owner),
						AssetId:      graphql_scalars.NewBytes32TryFromStringOrPanic(utxo.AssetId),
						Amount:       utxo.Amount,
						BlockCreated: utxo.BlockCreated,
						TxCreatedIdx: utxo.TxCreatedIdx,
					}
					res = append(res, []*graphql_object.CoinStruct{&coin})
				}
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
