package graphql_schemas

import "github.com/graphql-go/graphql"

type ConsensusParametersVersionType struct {
	SchemaFields EnumFields
}

//	pub enum ConsensusParametersVersion {
//		V1,
//	 }
type ConsensusParametersVersionStruct struct {
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

func ConsensusParametersVersion() *ConsensusParametersVersionType {
	enumConfig := graphql.EnumConfig{
		Name: "ConsensusParametersVersion",
		Values: graphql.EnumValueConfigMap{
			"V1": &graphql.EnumValueConfig{
				Value: 1,
			},
		},
	}
	enum := graphql.NewEnum(enumConfig)

	return &ConsensusParametersVersionType{
		SchemaFields: EnumFields{
			EnumConfig: &enumConfig,
			Enum:       enum,
		},
	}
}
