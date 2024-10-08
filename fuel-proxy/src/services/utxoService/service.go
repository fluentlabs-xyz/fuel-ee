package utxoService

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fluentlabs-xyz/fuel-ee/src/config"
	"github.com/fluentlabs-xyz/fuel-ee/src/helpers"
	"github.com/fluentlabs-xyz/fuel-ee/src/repo/utxoRepo"
	"github.com/fluentlabs-xyz/fuel-ee/src/types"
	"github.com/go-redis/redis/v8"
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
		//log.Infof("UtxoBGProcessing: next cycle in %d sec", timeoutSec)
		time.Sleep(time.Duration(timeoutSec) * time.Second)
		s.processCycle()
	}
	for {
		doBeforeCycle()
	}
}

func (s *Service) processCycle() {
	lastProcessedBlockNumber, err := s.utxoRepo.LastProcessedBlockNumber(context.Background())
	if err != nil {
		log.Printf("error when getting LastProcessedBlockNumber: %s", err)
		return
	}

	blockBatchSize := int64(100)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(lastProcessedBlockNumber + 1)),
	}
	query.FromBlock.Add(query.FromBlock, big.NewInt(blockBatchSize))

	logItems, err := s.ethClient.FilterLogs(context.Background(), query)
	if err != nil {
		log.Printf("error when filtering logs: %s", err)
		return
	}
	prevBlockNumber := uint64(0)
	if len(logItems) > 0 {
		prevBlockNumber = logItems[0].BlockNumber
	}
	for _, logItem := range logItems {
		log.Printf("new logItem: %+v", logItem)
		selector := logItem.Topics[0]

		if bytes.Equal(selector.Bytes(), types.FvmDepositSigBytesAligned32) {
			if err = s.processDeposit(&logItem); err != nil {
				log.Printf("error while processing deposit: %s", err)
				break
			}
		} else if bytes.Equal(selector.Bytes(), types.FvmWithdrawSigBytesAligned32) {
			if err = s.processWithdraw(&logItem); err != nil {
				log.Printf("error while processing withdraw: %s", err)
				break
			}
		} else {
			log.Printf("unprocessed log item: %+v", logItem)
		}

		if logItem.BlockNumber > prevBlockNumber {
			err = s.utxoRepo.SaveLastProcessedBlockNumber(context.Background(), prevBlockNumber)
			if err != nil {
				log.Printf("failed to save last processed block number: %s", err)
				break
			}
			prevBlockNumber = logItem.BlockNumber
		}
	}
}
func (s *Service) processDeposit(logItem *ethtypes.Log) error {
	if len(logItem.Data) != 32+32+32+32+32 {
		return errors.New(fmt.Sprintf("cannot process log item, because it's data len is %d. logItem: %+v", len(logItem.Data), logItem))
	}
	idx := 0
	recipientAddressHex := helpers.BytesToHexStringPrefixed(logItem.Data[idx : idx+32])
	idx += 32
	amountHex := helpers.BytesToHexNumberStringPrefixed(logItem.Data[idx : idx+32])
	idx += 32
	txIdHex := helpers.BytesToHexStringPrefixed(logItem.Data[idx : idx+32])
	idx += 32
	txOutputIndexHex := helpers.BytesToHexNumberStringPrefixed(logItem.Data[idx : idx+32])
	idx += 32
	assetIdHex := helpers.BytesToHexStringPrefixed(logItem.Data[idx : idx+32])

	amount, err := helpers.HexStringToBinInt(amountHex)
	if err != nil {
		return errors.New(fmt.Sprintf("failed converting hex amount '%s' to big int: %s", amountHex, err))
	}

	err = s.utxoRepo.SaveOne(
		context.Background(),
		utxoRepo.NewUtxoEntity(txIdHex, txOutputIndexHex, recipientAddressHex, assetIdHex, amount.Uint64(), 0, 0),
	)
	if err != nil {
		return errors.New(fmt.Sprintf("failed persisting record: %s", err))
	}

	return nil
}

func (s *Service) processWithdraw(logItem *ethtypes.Log) error {
	if len(logItem.Data) != 32+32+32 {
		return errors.New(fmt.Sprintf("cannot process log item, because it's data len is %d. logItem: %+v", len(logItem.Data), logItem))
	}
	idx := 0
	ownerHex := helpers.BytesToHexStringPrefixed(logItem.Data[idx : idx+32])
	idx += 32
	txIdHex := helpers.BytesToHexStringPrefixed(logItem.Data[idx : idx+32])
	idx += 32
	txOutputIndexHex := helpers.BytesToHexNumberStringPrefixed(logItem.Data[idx : idx+32])

	err := s.utxoRepo.DeleteByFields(
		context.Background(),
		ownerHex,
		txIdHex,
		txOutputIndexHex,
	)
	if err != nil {
		return errors.New(fmt.Sprintf("failed persisting record: %s", err))
	}

	return nil
}
