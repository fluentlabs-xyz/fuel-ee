package graphql_schemas

import "github.com/graphql-go/graphql"

type GasCostsType struct {
	SchemaFields SchemaFields
}

//		pub struct GasCosts {
//	   pub version: GasCostsVersion,
//	   pub add: U64,
//	   pub addi: U64,
//	   pub and: U64,
//	   pub andi: U64,
//	   pub bal: U64,
//	   pub bhei: U64,
//	   pub bhsh: U64,
//	   pub burn: U64,
//	   pub cb: U64,
//	   pub cfsi: U64,
//	   pub div: U64,
//	   pub divi: U64,
//	   pub eck1: U64,
//	   pub ecr1: U64,
//	   pub ed19: U64,
//	   pub eq: U64,
//	   pub exp: U64,
//	   pub expi: U64,
//	   pub flag: U64,
//	   pub gm: U64,
//	   pub gt: U64,
//	   pub gtf: U64,
//	   pub ji: U64,
//	   pub jmp: U64,
//	   pub jne: U64,
//	   pub jnei: U64,
//	   pub jnzi: U64,
//	   pub jmpf: U64,
//	   pub jmpb: U64,
//	   pub jnzf: U64,
//	   pub jnzb: U64,
//	   pub jnef: U64,
//	   pub jneb: U64,
//	   pub lb: U64,
//	   pub log: U64,
//	   pub lt: U64,
//	   pub lw: U64,
//	   pub mint: U64,
//	   pub mlog: U64,
//	   pub mod_op: U64,
//	   pub modi: U64,
//	   pub move_op: U64,
//	   pub movi: U64,
//	   pub mroo: U64,
//	   pub mul: U64,
//	   pub muli: U64,
//	   pub mldv: U64,
//	   pub noop: U64,
//	   pub not: U64,
//	   pub or: U64,
//	   pub ori: U64,
//	   pub poph: U64,
//	   pub popl: U64,
//	   pub pshh: U64,
//	   pub pshl: U64,
//	   pub ret: U64,
//	   pub rvrt: U64,
//	   pub sb: U64,
//	   pub sll: U64,
//	   pub slli: U64,
//	   pub srl: U64,
//	   pub srli: U64,
//	   pub srw: U64,
//	   pub sub: U64,
//	   pub subi: U64,
//	   pub sw: U64,
//	   pub sww: U64,
//	   pub time: U64,
//	   pub tr: U64,
//	   pub tro: U64,
//	   pub wdcm: U64,
//	   pub wqcm: U64,
//	   pub wdop: U64,
//	   pub wqop: U64,
//	   pub wdml: U64,
//	   pub wqml: U64,
//	   pub wddv: U64,
//	   pub wqdv: U64,
//	   pub wdmd: U64,
//	   pub wqmd: U64,
//	   pub wdam: U64,
//	   pub wqam: U64,
//	   pub wdmm: U64,
//	   pub wqmm: U64,
//	   pub xor: U64,
//	   pub xori: U64,
//
//	   pub aloc_dependent_cost: DependentCost,
//	   pub cfe: DependentCost,
//	   pub cfei_dependent_cost: DependentCost,
//	   pub call: DependentCost,
//	   pub ccp: DependentCost,
//	   pub croo: DependentCost,
//	   pub csiz: DependentCost,
//	   pub k256: DependentCost,
//	   pub ldc: DependentCost,
//	   pub logd: DependentCost,
//	   pub mcl: DependentCost,
//	   pub mcli: DependentCost,
//	   pub mcp: DependentCost,
//	   pub mcpi: DependentCost,
//	   pub meq: DependentCost,
//	   pub retd: DependentCost,
//	   pub s256: DependentCost,
//	   pub scwq: DependentCost,
//	   pub smo: DependentCost,
//	   pub srwq: DependentCost,
//	   pub swwq: DependentCost,
//
//	   // Non-opcodes prices
//	   pub contract_root: DependentCost,
//	   pub state_root: DependentCost,
//	   pub vm_initialization: DependentCost,
//	   pub new_storage_per_byte: U64,
//	}
type GasCostsStruct struct {
	aloc interface{}
}

func GasCosts(gasCostsVersionType *GasCostsVersionType, dependentCostType *DependentCostType) (*GasCostsType, error) {
	objectConfig := graphql.ObjectConfig{Name: "GasCosts", Fields: graphql.Fields{
		"version": &graphql.Field{
			Type: gasCostsVersionType.SchemaFields.Entity,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 1, nil
			},
		},
		// add
		"add": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// addi
		"addi": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// aloc
		"aloc": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// and
		"and": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// andi
		"andi": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// bal
		"bal": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 29, nil
			},
		},
		// bhei
		"bhei": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// bhsh
		"bhsh": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// burn
		"burn": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 19976, nil
			},
		},
		// cb
		"cb": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// cfei
		"cfei": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// cfsi
		"cfsi": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// div
		"div": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// divi
		"divi": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// ecr1
		"ecr1": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 26135, nil
			},
		},
		// eck1
		"eck1": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 1907, nil
			},
		},
		// ed19
		"ed19": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 1893, nil
			},
		},
		// eq
		"eq": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// exp
		"exp": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// expi
		"expi": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// flag
		"flag": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// gm
		"gm": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// gt
		"gt": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// gtf
		"gtf": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 13, nil
			},
		},
		// ji
		"ji": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// jmp
		"jmp": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// jne
		"jne": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// jnei
		"jnei": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// jnzi
		"jnzi": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// jmpf
		"jmpf": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// jmpb
		"jmpb": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// jnzf
		"jnzf": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// jnzb
		"jnzb": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// jnef
		"jnef": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// jneb
		"jneb": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// lb
		"lb": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// log
		"log": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 102, nil
			},
		},
		// lt
		"lt": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// lw
		"lw": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// mint
		"mint": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 18042, nil
			},
		},
		// mlog
		"mlog": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// modOp
		"modOp": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// modi
		"modi": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// moveOp
		"moveOp": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// movi
		"movi": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// mroo
		"mroo": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 4, nil
			},
		},
		// mul
		"mul": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// muli
		"muli": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// mldv
		"mldv": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 3, nil
			},
		},
		// noop
		"noop": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 1, nil
			},
		},
		// not
		"not": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// or
		"or": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// ori
		"ori": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// poph
		"poph": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 3, nil
			},
		},
		// popl
		"popl": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 3, nil
			},
		},
		// pshh
		"pshh": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 5, nil
			},
		},
		// pshl
		"pshl": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 5, nil
			},
		},
		// ret
		"ret": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 53, nil
			},
		},
		// rvrt
		"rvrt": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 52, nil
			},
		},
		// sb
		"sb": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// sll
		"sll": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// slli
		"slli": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// srl
		"srl": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// srli
		"srli": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// srw
		"srw": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 177, nil
			},
		},
		// sub
		"sub": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// subi
		"subi": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// sw
		"sw": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// sww
		"sww": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 17302, nil
			},
		},
		// time
		"time": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 35, nil
			},
		},
		// tr
		"tr": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 27852, nil
			},
		},
		// tro
		"tro": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 19718, nil
			},
		},
		// wdcm
		"wdcm": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// wqcm
		"wqcm": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// wdop
		"wdop": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 3, nil
			},
		},
		// wqop
		"wqop": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 3, nil
			},
		},
		// wdml
		"wdml": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 3, nil
			},
		},
		// wqml
		"wqml": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 3, nil
			},
		},
		// wddv
		"wddv": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 4, nil
			},
		},
		// wqdv
		"wqdv": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 5, nil
			},
		},
		// wdmd
		"wdmd": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 8, nil
			},
		},
		// wqmd
		"wqmd": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 12, nil
			},
		},
		// wdam
		"wdam": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 7, nil
			},
		},
		// wqam
		"wqam": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 8, nil
			},
		},
		// wdmm
		"wdmm": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 8, nil
			},
		},
		// wqmm
		"wqmm": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 8, nil
			},
		},
		// xor
		"xor": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// xori
		"xori": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 2, nil
			},
		},
		// newStoragePerByte
		"newStoragePerByte": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 63, nil
			},
		},
		//  alocDependentCost {
		//    ...DependentCostFragment
		//  }
		"alocDependentCost": &graphql.Field{
			Type: dependentCostType.SchemaFields.Entity,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &LightOperationStruct{
					Base:        2,
					UnitsPerGas: 15,
				}, nil
			},
		},
		//  cfe {
		//    ...DependentCostFragment
		//  }
		"cfe": &graphql.Field{
			Type: dependentCostType.SchemaFields.Entity,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &LightOperationStruct{
					Base:        10,
					UnitsPerGas: 1818181,
				}, nil
			},
		},
		//  cfeiDependentCost {
		//    ...DependentCostFragment
		//  }
		"cfeiDependentCost": &graphql.Field{
			Type: dependentCostType.SchemaFields.Entity,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &LightOperationStruct{
					Base:        2,
					UnitsPerGas: 1000000,
				}, nil
			},
		},
		//  call {
		//    ...DependentCostFragment
		//  }
		"call": &graphql.Field{
			Type: dependentCostType.SchemaFields.Entity,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &LightOperationStruct{
					Base:        13513,
					UnitsPerGas: 7,
				}, nil
			},
		},
		//  ccp {
		//    ...DependentCostFragment
		//  }
		"ccp": &graphql.Field{
			Type: dependentCostType.SchemaFields.Entity,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &LightOperationStruct{
					Base:        34,
					UnitsPerGas: 39,
				}, nil
			},
		},
		//  croo {
		//    ...DependentCostFragment
		//  }
		"croo": &graphql.Field{
			Type: dependentCostType.SchemaFields.Entity,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &LightOperationStruct{
					Base:        91,
					UnitsPerGas: 3,
				}, nil
			},
		},
		//  csiz {
		//    ...DependentCostFragment
		//  }
		"csiz": &graphql.Field{
			Type: dependentCostType.SchemaFields.Entity,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &LightOperationStruct{
					Base:        31,
					UnitsPerGas: 438,
				}, nil
			},
		},
		//  k256 {
		//    ...DependentCostFragment
		//  }
		"k256": &graphql.Field{
			Type: dependentCostType.SchemaFields.Entity,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &LightOperationStruct{
					Base:        27,
					UnitsPerGas: 5,
				}, nil
			},
		},
		//  ldc {
		//    ...DependentCostFragment
		//  }
		"ldc": &graphql.Field{
			Type: dependentCostType.SchemaFields.Entity,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &LightOperationStruct{
					Base:        43,
					UnitsPerGas: 102,
				}, nil
			},
		},
		//  logd {
		//    ...DependentCostFragment
		//  }
		"logd": &graphql.Field{
			Type: dependentCostType.SchemaFields.Entity,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &LightOperationStruct{
					Base:        363,
					UnitsPerGas: 4,
				}, nil
			},
		},
		//  mcl {
		//    ...DependentCostFragment
		//  }
		"mcl": &graphql.Field{
			Type: dependentCostType.SchemaFields.Entity,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &LightOperationStruct{
					Base:        2,
					UnitsPerGas: 1041,
				}, nil
			},
		},
		//  mcli {
		//    ...DependentCostFragment
		//  }
		"mcli": &graphql.Field{
			Type: dependentCostType.SchemaFields.Entity,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &LightOperationStruct{
					Base:        2,
					UnitsPerGas: 1025,
				}, nil
			},
		},
		//  mcp {
		//    ...DependentCostFragment
		//  }
		"mcp": &graphql.Field{
			Type: dependentCostType.SchemaFields.Entity,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &LightOperationStruct{
					Base:        4,
					UnitsPerGas: 325,
				}, nil
			},
		},
		//  mcpi {
		//    ...DependentCostFragment
		//  }
		"mcpi": &graphql.Field{
			Type: dependentCostType.SchemaFields.Entity,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &LightOperationStruct{
					Base:        8,
					UnitsPerGas: 511,
				}, nil
			},
		},
		//  meq {
		//    ...DependentCostFragment
		//  }
		"meq": &graphql.Field{
			Type: dependentCostType.SchemaFields.Entity,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &LightOperationStruct{
					Base:        3,
					UnitsPerGas: 940,
				}, nil
			},
		},
		//  retd {
		//    ...DependentCostFragment
		//  }
		"retd": &graphql.Field{
			Type: dependentCostType.SchemaFields.Entity,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &LightOperationStruct{
					Base:        305,
					UnitsPerGas: 4,
				}, nil
			},
		},
		//  s256 {
		//    ...DependentCostFragment
		//  }
		"s256": &graphql.Field{
			Type: dependentCostType.SchemaFields.Entity,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &LightOperationStruct{
					Base:        31,
					UnitsPerGas: 4,
				}, nil
			},
		},
		//  scwq {
		//    ...DependentCostFragment
		//  }
		"scwq": &graphql.Field{
			Type: dependentCostType.SchemaFields.Entity,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &HeavyOperationStruct{
					Base:       16346,
					GasPerUnit: 17163,
				}, nil
			},
		},
		//  smo {
		//    ...DependentCostFragment
		//  }
		"smo": &graphql.Field{
			Type: dependentCostType.SchemaFields.Entity,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &LightOperationStruct{
					Base:        40860,
					UnitsPerGas: 2,
				}, nil
			},
		},
		//  srwq {
		//    ...DependentCostFragment
		//  }
		"srwq": &graphql.Field{
			Type: dependentCostType.SchemaFields.Entity,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &HeavyOperationStruct{
					Base:       187,
					GasPerUnit: 179,
				}, nil
			},
		},
		//  swwq {
		//    ...DependentCostFragment
		//  }
		"swwq": &graphql.Field{
			Type: dependentCostType.SchemaFields.Entity,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &HeavyOperationStruct{
					Base:       17046,
					GasPerUnit: 16232,
				}, nil
			},
		},
		//  contractRoot {
		//    ...DependentCostFragment
		//  }
		"contractRoot": &graphql.Field{
			Type: dependentCostType.SchemaFields.Entity,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &LightOperationStruct{
					Base:        31,
					UnitsPerGas: 2,
				}, nil
			},
		},
		//  stateRoot {
		//    ...DependentCostFragment
		//  }
		"stateRoot": &graphql.Field{
			Type: dependentCostType.SchemaFields.Entity,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &HeavyOperationStruct{
					Base:       236,
					GasPerUnit: 122,
				}, nil
			},
		},
		//  vmInitialization {
		//    ...DependentCostFragment
		//  }
		"vmInitialization": &graphql.Field{
			Type: dependentCostType.SchemaFields.Entity,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &LightOperationStruct{
					Base:        3957,
					UnitsPerGas: 48,
				}, nil
			},
		},
		/*"aloc": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return 1, nil
			},
		},
		"alocDependentCost": &graphql.Field{
			Type: dependentCostType.SchemaFields.Entity,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &LightOperationStruct{
					Base:        1,
					UnitsPerGas: 1,
				}, nil
			},
		},*/
	}}
	object := graphql.NewObject(objectConfig)
	schemaConfig := graphql.SchemaConfig{Query: object}
	schema, err := graphql.NewSchema(schemaConfig)

	return &GasCostsType{
		SchemaFields: SchemaFields{
			Schema:       &schema,
			ObjectConfig: &objectConfig,
			Object:       object,
			SchemaConfig: &schemaConfig,
		},
	}, err
}
