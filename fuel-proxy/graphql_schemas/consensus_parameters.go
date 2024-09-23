package graphql_schemas

import "github.com/graphql-go/graphql"

type ConsensusParametersType struct {
	SchemaFields SchemaFields
}

//	pub struct ConsensusParameters {
//	   pub version: ConsensusParametersVersion,
//	   pub tx_params: TxParameters,
//	   pub predicate_params: PredicateParameters,
//	   pub script_params: ScriptParameters,
//	   pub contract_params: ContractParameters,
//	   pub fee_params: FeeParameters,
//	   pub base_asset_id: AssetId,
//	   pub block_gas_limit: U64,
//	   pub chain_id: U64,
//	   pub gas_costs: GasCosts,
//	   pub privileged_address: Address,
//	}
type ConsensusParametersStruct struct {
	// pub version: ConsensusParametersVersion,
	// pub tx_params: TxParameters,
	// pub predicate_params: PredicateParameters,
	// pub script_params: ScriptParameters,
	// pub contract_params: ContractParameters,
	// pub fee_params: FeeParameters,
	// pub base_asset_id: AssetId,
	// pub block_gas_limit: U64,
	// pub chain_id: U64,
	// pub gas_costs: GasCosts,
	// pub privileged_address: Address,
}

func ConsensusParameters(consensusParametersVersionType *ConsensusParametersVersionType) (*ConsensusParametersType, error) {

	objectConfig := graphql.ObjectConfig{Name: "ConsensusParameters", Fields: graphql.Fields{
		"chainId": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 0, nil
			},
		},
		"version": &graphql.Field{
			Type: consensusParametersVersionType.SchemaFields.Enum,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 1, nil
			},
		},
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	return &ConsensusParametersType{
		SchemaFields: SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
