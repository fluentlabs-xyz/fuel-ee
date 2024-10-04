package graphqlServerService

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fluentlabs-xyz/fuel-ee/src/config"
	"github.com/fluentlabs-xyz/fuel-ee/src/graphql_entrypoints"
	"github.com/fluentlabs-xyz/fuel-ee/src/graphql_object"
	"github.com/fluentlabs-xyz/fuel-ee/src/helpers"
	"github.com/fluentlabs-xyz/fuel-ee/src/services/utxoService"
	"github.com/go-redis/redis/v8"
	"github.com/graphql-go/graphql"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Service struct {
	config    *config.Config
	redis     *redis.Client
	ethClient *ethclient.Client
}

func New(
	config *config.Config,
	redis *redis.Client,
	ethClient *ethclient.Client,
	utxoService *utxoService.Service,
) *Service {
	consensusParametersVersionType := graphql_object.ConsensusParametersVersion()

	lightOperationType, err := graphql_object.LightOperation()
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	heavyOperationType, err := graphql_object.HeavyOperation()
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	dependentCostType := graphql_object.DependentCost(lightOperationType, heavyOperationType)

	contractParametersVersionType := graphql_object.ContractParametersVersion()
	contractParametersType, err := graphql_object.ContractParameters(contractParametersVersionType)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	gasCostsVersionType := graphql_object.GasCostsVersion()
	gasCostsType, err := graphql_object.GasCosts(gasCostsVersionType, dependentCostType)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	feeParametersVersionType := graphql_object.FeeParametersVersion()
	feeParametersType, err := graphql_object.FeeParameters(feeParametersVersionType)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	scriptParametersVersionType := graphql_object.ScriptParametersVersion()
	scriptParametersType, err := graphql_object.ScriptParameters(scriptParametersVersionType)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	predicateParametersVersionType := graphql_object.PredicateParametersVersion()
	predicateParametersType, err := graphql_object.PredicateParameters(predicateParametersVersionType)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	txParametersVersionType := graphql_object.TxParametersVersion()
	txParametersType, err := graphql_object.TxParameters(txParametersVersionType)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	consensusParametersType, err := graphql_object.ConsensusParameters(
		consensusParametersVersionType,
		txParametersType,
		predicateParametersType,
		scriptParametersType,
		contractParametersType,
		feeParametersType,
		gasCostsType,
	)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	transactionType, err := graphql_object.Transaction()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	headerType, err := graphql_object.Header()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	blockType, err := graphql_object.Block(headerType, transactionType)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	chainInfoType, err := graphql_object.ChainInfo(blockType, consensusParametersType)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	nodeInfoType, err := graphql_object.NodeInfo()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	pageInfoType, err := graphql_object.PageInfo()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	gasPriceType, err := graphql_object.GasPrice()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	programStateType, err := graphql_object.ProgramState()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	dryRunSuccessStatusType, err := graphql_object.MakeDryRunSuccessStatus(programStateType)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	dryRunFailureStatusType, err := graphql_object.MakeDryRunFailureStatus(programStateType)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	dryRunTransactionStatusType := graphql_object.MakeDryRunTransactionStatus(dryRunSuccessStatusType, dryRunFailureStatusType)

	receiptTypeType := graphql_object.MakeReceiptType()

	receiptType, err := graphql_object.Receipt(receiptTypeType)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	dryRunTransactionExecutionStatusType, err := graphql_object.DryRunTransactionExecutionStatus(dryRunTransactionStatusType, receiptType)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	submitType, err := graphql_object.MakeSubmitType()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// ENTRIES

	getChainEntry, err := graphql_entrypoints.MakeGetChainEntry(chainInfoType)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	getNodeInfoEntry, err := graphql_entrypoints.MakeGetNodeInfoEntry(nodeInfoType)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	getCoinsEntry, err := graphql_entrypoints.MakeGetCoinsEntry(pageInfoType)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	estimateGasPriceEntry, err := graphql_entrypoints.MakeEstimateGasPriceEntry(gasPriceType)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	dryRunEntry, err := graphql_entrypoints.MakeDryRunEntry(ethClient, dryRunTransactionExecutionStatusType)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	submitEntry, err := graphql_entrypoints.MakeSubmitEntry(ethClient, submitType)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	coinType, err := graphql_object.MakeCoin()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	messageCoinType, err := graphql_object.MakeMessageCoin()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	coinTypeType := graphql_object.MakeCoinType(coinType, messageCoinType)

	getCoinsToSpendEntry, err := graphql_entrypoints.MakeGetCoinsToSpendEntry(utxoService, coinTypeType)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/v1/graphql", func(w http.ResponseWriter, req *http.Request) {
		var p postData
		if err := json.NewDecoder(req.Body).Decode(&p); err != nil {
			errText := fmt.Sprintf("failed to decode: %s", err)
			log.Printf(errText)
			_, _ = helpers.HttpWriteError(w, http.StatusBadRequest, errText)
			return
		}
		log.Printf("Operation: '%s' Variables: '%s' Query: '%s'", p.Operation, p.Variables, p.Query)
		var schema *graphql.Schema

		switch p.Operation {
		case "getChain":
			schema = getChainEntry.SchemaFields.Schema
		case "getNodeInfo":
			schema = getNodeInfoEntry.SchemaFields.Schema
		case "getCoins":
			_, _ = helpers.HttpWriteError(w, http.StatusMethodNotAllowed, "temporarily disabled")
			return
			schema = getCoinsEntry.SchemaFields.Schema
		case "estimateGasPrice":
			schema = estimateGasPriceEntry.SchemaFields.Schema
		case "dryRun":
			schema = dryRunEntry.SchemaFields.Schema
		case "submit":
			schema = submitEntry.SchemaFields.Schema
		case "getCoinsToSpend":
			schema = getCoinsToSpendEntry.SchemaFields.Schema
		default:
			errText := fmt.Sprintf("unsupported operation: %s", p.Operation)
			log.Printf(errText)
			_, _ = helpers.HttpWriteError(w, http.StatusBadRequest, errText)
			return
		}
		params := graphql.Params{
			Context:        req.Context(),
			Schema:         *schema,
			RequestString:  p.Query,
			VariableValues: p.Variables,
			OperationName:  p.Operation,
		}
		result := graphql.Do(params)
		if len(result.Errors) > 0 {
			log.Printf("graphql errors: %s", result.Errors)
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(result); err != nil {
			log.Printf("could not write result to response: %s", err)
		}
	})

	return &Service{
		config: config,
		redis:  redis,
	}
}

func (s *Service) Start() error {
	go s.startServer()
	return nil
}

func (s *Service) Stop() error {
	return nil
}

func (s *Service) startServer() {
	port := s.config.GraphQL.Port
	log.Printf("Server is running on port %d", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}
