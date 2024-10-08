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
	"strings"
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

const filterArgName = "filter"
const afterArgName = "after"
const beforeArgName = "before"
const firstArgName = "first"
const lastArgName = "last"

func MakeGetCoinsEntry(utxoService *utxoService.Service, getCoinsType *graphql_object.GetCoinsType) (*GetCoinsEntry, error) {
	objectConfig := graphql.ObjectConfig{Name: "GetCoinsEntry", Fields: graphql.Fields{
		"coins": &graphql.Field{
			Type: getCoinsType.SchemaFields.Object,
			Args: graphql.FieldConfigArgument{
				filterArgName: &graphql.ArgumentConfig{
					Type: graphql_input_objects.CoinFilterInput,
				},
				afterArgName: &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				beforeArgName: &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				firstArgName: &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				lastArgName: &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				filterArg := p.Args[filterArgName]
				afterArg := p.Args[afterArgName]
				beforeArg := p.Args[beforeArgName]
				firstArg := p.Args[firstArgName]
				lastArg := p.Args[lastArgName]
				log.Printf(
					"filterArg '%+v' afterArg '%+v' beforeArg '%+v' firstArg '%+v' lastArg '%+v' ",
					filterArg,
					afterArg,
					beforeArg,
					firstArg,
					lastArg,
				)
				var owner *graphql_scalars.Bytes32
				var assetId *graphql_scalars.Bytes32
				if filterArg != nil {
					filterArgMap, ok := filterArg.(map[string]interface{})
					if !ok {
						errText := "param [filter] must be of map[string]interface{} type"
						log.Printf(errText)
						return nil, errors.New(errText)
					}
					vInterface, ok := filterArgMap["owner"]
					if !ok {
						errText := fmt.Sprintf("field [owner] is a mandatory")
						log.Printf(errText)
						return nil, errors.New(errText)
					}
					owner, ok = vInterface.(*graphql_scalars.Bytes32)
					if !ok {
						errText := fmt.Sprintf("failed to convert [owner] param")
						log.Printf(errText)
						return nil, errors.New(errText)
					}
					vInterface, ok = filterArgMap["assetId"]
					if ok {
						assetId, ok = vInterface.(*graphql_scalars.Bytes32)
						if !ok {
							errText := fmt.Sprintf("failed to convert [assetId] param")
							log.Printf(errText)
							return nil, errors.New(errText)
						}
					}
				}
				// TODO query coins here
				utxos, err := utxoService.Repo().FindAllByParams(context.Background(), strings.ToLower(owner.String()), "*", "*", false)
				if err != nil {
					errText := "param [excludedIdsArgInterfaces] must have 'utxos' key which values are of []interface{} type"
					log.Printf(errText)
					return nil, errors.New(errText)
				}
				if assetId != nil {
					for _, utxo := range utxos {
						if utxo.AssetId != assetId.String() {
							delete(utxos, assetId.String())
						}
					}
				}

				edges := make([]*graphql_object.NodeStruct, 0, len(utxos))
				for _, utxo := range utxos {
					assetId, err = utxo.GetAssetId()
					if err != nil {
						return nil, errors.New(fmt.Sprintf("failed to get asset id: %s", err))
					}
					owner, err = utxo.GetOwner()
					if err != nil {
						return nil, errors.New(fmt.Sprintf("failed to get asset id: %s", err))
					}
					utxoId, err := utxo.UtxoId()
					if err != nil {
						return nil, errors.New(fmt.Sprintf("failed to get asset id: %s", err))
					}
					edges = append(edges, &graphql_object.NodeStruct{
						Node: &graphql_object.CoinStruct{
							Amount:       utxo.Amount,
							BlockCreated: utxo.BlockCreated,
							TxCreatedIdx: utxo.TxCreatedIdx,
							AssetId:      assetId,
							Owner:        owner,
							UtxoId:       utxoId,
						},
					})
				}

				return &graphql_object.GetCoinsStruct{
					PageInfo: &graphql_object.PageInfoStruct{},
					Edges:    edges,
				}, nil
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
