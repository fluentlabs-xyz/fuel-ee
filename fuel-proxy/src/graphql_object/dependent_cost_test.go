package graphql_object

import (
	"encoding/json"
	"github.com/graphql-go/graphql"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestDependentCost(t *testing.T) {
	query := `
	 {
	   fff {
	     ...DependentCostFragment
	   }
	 }
		
	fragment DependentCostFragment on DependentCost {
	  ... on LightOperation {
		type: __typename
		base
		unitsPerGas
	  }
	  ... on HeavyOperation {
		type: __typename
		base
		gasPerUnit
	  }
	}
    `

	lightOperationType, err := LightOperation()
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	heavyOperationType, err := HeavyOperation()
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	dependentCostType := DependentCost(lightOperationType, heavyOperationType)

	objectConfig := graphql.ObjectConfig{Name: "SomeRoot", Fields: graphql.Fields{
		"fff": &graphql.Field{
			Type: dependentCostType.SchemaFields.Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &LightOperationStruct{
					Base:        1,
					UnitsPerGas: 2,
				}, nil
			},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	log.Printf("test response: %s", rJSON)
}
