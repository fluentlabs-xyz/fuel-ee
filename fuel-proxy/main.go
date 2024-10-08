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

	helpers.RequireNoError(di.Provide(func(config *config.Config) (*ethclient.Client, error) {
		return ethclient.Dial(config.EthProvider.Url)
	}))

	helpers.RequireNoError(di.Provide(utxoService.New))
	helpers.RequireNoError(di.Provide(graphqlServerService.New))

	helpers.RequireNoError(di.Invoke(func(utxoService *utxoService.Service, ethClient *ethclient.Client) {
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
