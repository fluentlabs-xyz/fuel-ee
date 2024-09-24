package graphql_schemas

import "github.com/graphql-go/graphql"

type NodeInfoType struct {
	SchemaFields SchemaFields
}

//	pub struct NodeInfo {
//	   pub utxo_validation: bool,
//	   pub vm_backtrace: bool,
//	   pub max_tx: U64,
//	   pub max_depth: U64,
//	   pub node_version: String,
//	}
//
// {"data":{"nodeInfo":{"utxoValidation":true,"vmBacktrace":false,"maxTx":"4064","maxDepth":"10","nodeVersion":"0.31.0"}}}
type NodeInfoStruct struct {
	UtxoValidation bool   `json:"utxoValidation"`
	VmBacktrace    bool   `json:"vmBacktrace"`
	MaxTx          string `json:"maxTx"`
	MaxDepth       string `json:"maxDepth"`
	NodeVersion    string `json:"nodeVersion"`
}

func NodeInfo() (*NodeInfoType, error) {
	objectConfig := graphql.ObjectConfig{Name: "NodeInfo", Fields: graphql.Fields{
		"utxoValidation": &graphql.Field{
			Type: graphql.Boolean,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return true, nil
			},
		},
		"vmBacktrace": &graphql.Field{
			Type: graphql.Boolean,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return false, nil
			},
		},
		"maxTx": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 4064, nil
			},
		},
		"maxDepth": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 10, nil
			},
		},
		"nodeVersion": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "0.31.0", nil
			},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	return &NodeInfoType{
		SchemaFields: SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
