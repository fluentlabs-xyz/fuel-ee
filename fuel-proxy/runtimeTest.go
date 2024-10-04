package main

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fluentlabs-xyz/fuel-ee/src/services/utxoService"
	"log"
)

func runTest(ethClient *ethclient.Client, utxoService *utxoService.Service) {
	res, err := utxoService.Repo().FindAllByParams(
		context.Background(),
		"*",
		"*",
		"*",
		false,
	)
	if err != nil {
		log.Printf("error when FindAllByParams: %s", err)
	} else {
		log.Printf("found %d utxos", len(res))
		//for key, utxo := range res {
		//	log.Printf("for key '%s' utxo '%+v'", key, utxo)
		//}
	}
}
