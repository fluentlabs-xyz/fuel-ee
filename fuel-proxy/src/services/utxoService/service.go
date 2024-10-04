package utxoService

import (
	"bytes"
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fluentlabs-xyz/fuel-ee/src/config"
	"github.com/fluentlabs-xyz/fuel-ee/src/helpers"
	"github.com/fluentlabs-xyz/fuel-ee/src/repo/utxoRepo"
	"github.com/fluentlabs-xyz/fuel-ee/src/types"
	"github.com/go-redis/redis/v8"
	"github.com/holiman/uint256"
	log "github.com/sirupsen/logrus"
	"math/big"
	"time"
)

type Service struct {
	config      *config.Config
	redisClient *redis.Client
	ethClient   *ethclient.Client
	utxoRepo    *utxoRepo.UtxoRepo
}

func New(
	config *config.Config,
	redisClient *redis.Client,
	ethClient *ethclient.Client,
) *Service {
	return &Service{
		config:      config,
		redisClient: redisClient,
		ethClient:   ethClient,
		utxoRepo:    utxoRepo.NewUtxoRepo(redisClient),
	}
}

func (s *Service) Start() error {
	go s.startBgProcessing()
	return nil
}

func (s *Service) Stop() error {
	return nil
}

func (s *Service) Repo() *utxoRepo.UtxoRepo {
	return s.utxoRepo
}

func (s *Service) startBgProcessing() {
	timeoutSec := s.config.App.UtxoBGProcessingTimeoutSec
	doBeforeCycle := func() {
		log.Infof("UtxoBGProcessing: next cycle in %d sec", timeoutSec)
		time.Sleep(time.Duration(timeoutSec) * time.Second)
	}
	for {
		doBeforeCycle()

		lastProcessedBlockNumber, err := s.utxoRepo.LastProcessedBlockNumber(context.Background())
		if err != nil {
			log.Printf("error when getting LastProcessedBlockNumber: %s", err)
		}

		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(lastProcessedBlockNumber + 1)),
		}
		logItems, err := s.ethClient.FilterLogs(context.Background(), query)
		if err != nil {
			log.Fatal(err)
		}
		for _, logItem := range logItems {
			log.Printf("found new logItem: %+v", logItem)
			if len(logItem.Data) != 160 {
				log.Printf("error: cannot add log item, because it's data len is %d. logItem: %+v. processing stopped", len(logItem.Data), logItem)
				break
			}
			amountHex := helpers.BytesToHexNumberStringPrefixed(logItem.Data[:32])
			txIdHex := helpers.BytesToHexStringPrefixed(logItem.Data[32:64])
			txOutputIndexHex := helpers.BytesToHexNumberStringPrefixed(logItem.Data[64:96])
			recipientAddressHex := helpers.BytesToHexStringPrefixed(logItem.Data[96:128])
			assetIdHex := helpers.BytesToHexStringPrefixed(logItem.Data[128:160])
			selector := logItem.Topics[0]

			// filter deposits only
			if !bytes.Equal(selector.Bytes(), types.FvmDepositSig32Bytes) {
				continue
			}
			amount, err := uint256.FromHex(amountHex)
			if err != nil {
				log.Printf("when converting amount '%s' got error: %s. processing stopped", amountHex, err)
				break
			}

			err = s.utxoRepo.SaveOne(
				context.Background(),
				utxoRepo.NewUtxoEntity(txIdHex, txOutputIndexHex, recipientAddressHex, assetIdHex, amount.Uint64(), 1, 0),
			)
			if err != nil {
				log.Printf("failed to process logs, error: %s. processing stopped", err)
				break
			}

			err = s.utxoRepo.SaveLastProcessedBlockNumber(context.Background(), logItem.BlockNumber)
			if err != nil {
				log.Printf("error when saving last processed block number: %s", err)
			}
		}
	}
}
