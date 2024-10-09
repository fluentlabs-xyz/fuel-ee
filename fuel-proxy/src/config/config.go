package config

import (
	"github.com/fluentlabs-xyz/fuel-ee/src/helpers"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

type AppConfig struct {
	UtxoBGProcessingTimeoutSec int64
}

func (c *AppConfig) parse() {
	c.UtxoBGProcessingTimeoutSec = viperGetOrDefaultInt("app.utxo-bg-processing-timeout-sec", 2)
}

type BlockchainConfig struct {
	ChainId               int64
	FvmDepositSig         uint32
	FvmDepositSigBytes    []byte
	FvmDepositSigBytes32  []byte
	FvmWithdrawSig        uint32
	FvmWithdrawSigBytes   []byte
	FvmWithdrawSigBytes32 []byte
	FvmDryRunSig          uint32
	FvmDryRunSigBytes     []byte
	FvmDryRunSigBytes32   []byte
	FvmExecSig            uint32
	FvmExecSigBytes       []byte
	FvmExecSigBytes32     []byte
}

func (c *BlockchainConfig) parse() {
	c.ChainId = viperGetOrDefaultInt("app.chain-id", 1337)

	c.FvmDepositSig = viperGetOrDefaultUint32("app.fvm-deposit-sig", 2311579102)
	c.FvmDepositSigBytes = helpers.Uint32ToBytesBEMust(c.FvmDepositSig, 4)
	c.FvmDepositSigBytes32 = helpers.Uint32ToBytesBEMust(c.FvmDepositSig, 32)

	c.FvmWithdrawSig = viperGetOrDefaultUint32("app.fvm-withdraw-sig", 3481020119)
	c.FvmWithdrawSigBytes = helpers.Uint32ToBytesBEMust(c.FvmWithdrawSig, 4)
	c.FvmWithdrawSigBytes32 = helpers.Uint32ToBytesBEMust(c.FvmWithdrawSig, 32)

	c.FvmDryRunSig = viperGetOrDefaultUint32("app.fvm-dry-run-sig", 1492973092)
	c.FvmDryRunSigBytes = helpers.Uint32ToBytesBEMust(c.FvmDryRunSig, 4)
	c.FvmDryRunSigBytes32 = helpers.Uint32ToBytesBEMust(c.FvmDryRunSig, 32)

	c.FvmExecSig = viperGetOrDefaultUint32("app.fvm-exec-sig", 3011009299)
	c.FvmExecSigBytes = helpers.Uint32ToBytesBEMust(c.FvmExecSig, 4)
	c.FvmExecSigBytes32 = helpers.Uint32ToBytesBEMust(c.FvmExecSig, 32)
}

type RedisConfig struct {
	Address  string
	User     string
	Password string
}

func (c *RedisConfig) parse() {
	c.Address = viperGetOrDefaultString("redis.address", "localhost:6379")
	c.User = viperGetOrDefaultString("redis.user", "")
	c.Password = viperGetOrDefaultString("redis.password", "123456")
}

type RelayerConfig struct {
	PrivateKey string
}

func (c *RelayerConfig) parse() {
	c.PrivateKey = viperGetOrDefaultString("relayer.private_key", "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
}

type EthProviderConfig struct {
	Url string
}

func (c *EthProviderConfig) parse() {
	c.Url = viperGetOrDefaultString("eth_provider.url", "http://127.0.0.1:8545")
}

type GraphQlConfig struct {
	Port int64
}

func (c *GraphQlConfig) parse() {
	c.Port = viperGetOrDefaultInt("graphql.port", 8080)
}

type Config struct {
	App         AppConfig
	Redis       RedisConfig
	Blockchain  BlockchainConfig
	GraphQL     GraphQlConfig
	EthProvider EthProviderConfig
	Relayer     RelayerConfig
}

func NewConfig() *Config {
	config := &Config{}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.SetDefault("config-path", "./config.yaml")
	configPath := viper.GetString("config-path")
	if _, err := os.Stat(configPath); err == nil {
		viper.SetConfigType(filepath.Ext(configPath)[1:])
		viper.SetConfigFile(configPath)
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("failed to read config (%s): %+v", configPath, err)
		}
	}

	config.Redis.parse()
	config.GraphQL.parse()
	config.EthProvider.parse()
	config.Relayer.parse()
	config.Blockchain.parse()
	config.App.parse()

	return config
}
