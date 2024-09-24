package main

import (
	"encoding/json"
	"fmt"
	"github.com/fluentlabs-xyz/fuel-ee/graphql_schemas"
	"github.com/graphql-go/graphql"
	"log"
	"net/http"
)

type postData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operationName"`
	Variables map[string]interface{} `json:"variables"`
}

func main() {
	//baseAssetIdObjectConfig := graphql.ObjectConfig{Name: "BaseAssetId", Fields: graphql.Fields{
	//	"id": &graphql.Field{
	//		Type: graphql.String,
	//		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
	//			return "0xf8f8b6283d7fa5b672b530cbb84fcccb4ff8dc40f8176ef4544ddb1f1952ad07", nil
	//		},
	//	},
	//}}
	//baseAssetIdObject := graphql.NewObject(baseAssetIdObjectConfig)
	//baseAssetIdSchemaConfig := graphql.SchemaConfig{Query: baseAssetIdObject}
	//_, err := graphql.NewSchema(baseAssetIdSchemaConfig)

	consensusParametersVersionType := graphql_schemas.ConsensusParametersVersion()

	lightOperationType, err := graphql_schemas.LightOperation()
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	heavyOperationType, err := graphql_schemas.HeavyOperation()
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	dependentCostType := graphql_schemas.DependentCost(lightOperationType, heavyOperationType)

	contractParametersVersionType := graphql_schemas.ContractParametersVersion()
	contractParametersType, err := graphql_schemas.ContractParameters(contractParametersVersionType)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	gasCostsVersionType := graphql_schemas.GasCostsVersion()
	gasCostsType, err := graphql_schemas.GasCosts(gasCostsVersionType, dependentCostType)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	feeParametersVersionType := graphql_schemas.FeeParametersVersion()
	feeParametersType, err := graphql_schemas.FeeParameters(feeParametersVersionType)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	scriptParametersVersionType := graphql_schemas.ScriptParametersVersion()
	scriptParametersType, err := graphql_schemas.ScriptParameters(scriptParametersVersionType)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	predicateParametersVersionType := graphql_schemas.PredicateParametersVersion()
	predicateParametersType, err := graphql_schemas.PredicateParameters(predicateParametersVersionType)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	txParametersVersionType := graphql_schemas.TxParametersVersion()
	txParametersType, err := graphql_schemas.TxParameters(txParametersVersionType)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	consensusParametersType, err := graphql_schemas.ConsensusParameters(
		consensusParametersVersionType,
		txParametersType,
		predicateParametersType,
		scriptParametersType,
		contractParametersType,
		feeParametersType,
		gasCostsType,
	)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	transactionType, err := graphql_schemas.Transaction()
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	headerType, err := graphql_schemas.Header()
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	blockType, err := graphql_schemas.Block(headerType, transactionType)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	chainInfoType, err := graphql_schemas.ChainInfo(blockType, consensusParametersType)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	chainType, err := graphql_schemas.Chain(chainInfoType)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	//// Query
	//query := `
	//	query getChain {
	//      ...chainInfoFragment
	//	}
	//	fragment chainInfoFragment on ChainInfo {
	//	  name
	//	  daHeight
	//      latestBlock {
	//	    ...blockFragment
	//      }
	//	}
	//	fragment blockFragment on Block {
	//	  id
	//	  height
	//	}
	//`
	//params := graphql.Params{Schema: *chainInfoType.SchemaFields.Schema, RequestString: query}
	//r := graphql.Do(params)
	//if len(r.Errors) > 0 {
	//	log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	//}
	//rJSON, _ := json.Marshal(r)
	//log.Printf("test response: %s", rJSON) // {“data”:{“hello”:”world”}}

	http.HandleFunc("/v1/graphql", func(w http.ResponseWriter, req *http.Request) {
		var p postData
		if err := json.NewDecoder(req.Body).Decode(&p); err != nil {
			log.Printf("failed to decode: %s", err)
			w.WriteHeader(400)
			return
		}
		log.Printf("query: %s", p.Query)
		params := graphql.Params{
			Context:        req.Context(),
			Schema:         *chainType.SchemaFields.Schema,
			RequestString:  p.Query,
			VariableValues: p.Variables,
			OperationName:  p.Operation,
		}
		result := graphql.Do(params)
		v, _ := json.Marshal(result)
		log.Printf("v: %s", v)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(result); err != nil {
			log.Printf("could not write result to response: %s", err)
		}
	})

	port := 8080
	log.Printf("Server is running on port %d", port)
	if err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}
