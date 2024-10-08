package config

import (
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
	Redis       RedisConfig
	GraphQL     GraphQlConfig
	EthProvider EthProviderConfig
	Relayer     RelayerConfig
	App         AppConfig
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
	config.App.parse()

	return config
}
