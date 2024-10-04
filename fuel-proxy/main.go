package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fluentlabs-xyz/fuel-ee/src/common"
	"github.com/fluentlabs-xyz/fuel-ee/src/config"
	"github.com/fluentlabs-xyz/fuel-ee/src/container"
	"github.com/fluentlabs-xyz/fuel-ee/src/helpers"
	"github.com/fluentlabs-xyz/fuel-ee/src/services/graphqlServerService"
	"github.com/fluentlabs-xyz/fuel-ee/src/services/utxoService"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
)

type postData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operationName"`
	Variables map[string]interface{} `json:"variables"`
}

func main() {
	helpers.RequireNoError(os.Setenv("TZ", "UTC"))
	di := container.CreateContainer()

	helpers.RequireNoError(
		di.Provide(
			func(config *config.Config) *redis.Client {
				redisOpts, err := redis.ParseURL(
					fmt.Sprintf(
						"redis://%s:%s@%s/",
						config.Redis.User,
						config.Redis.Password,
						config.Redis.Address,
					),
				)
				if err != nil {
					log.Fatalf("unable to parse redis URL: %s", err)
				}
				return redis.NewClient(redisOpts)
			},
		),
	)

	ethClient, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	helpers.RequireNoError(di.Provide(func() *ethclient.Client { return ethClient }))

	helpers.RequireNoError(di.Provide(utxoService.New))
	helpers.RequireNoError(di.Provide(graphqlServerService.New))

	helpers.RequireNoError(di.Invoke(func(utxoService *utxoService.Service) {
		runTest(ethClient, utxoService)
	}))

	container.MustInvoke(
		di, func(
			utxoService *utxoService.Service,
			graphqlServerService *graphqlServerService.Service,
		) {
			helpers.RequireNoError(utxoService.Start())
			helpers.RequireNoError(graphqlServerService.Start())

			common.WaitForSignal()

			helpers.RequireNoError(graphqlServerService.Stop())
			helpers.RequireNoError(utxoService.Stop())
		},
	)
}
