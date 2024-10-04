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

type GraphQlConfig struct {
	Port int64
}

func (c *GraphQlConfig) parse() {
	c.Port = viperGetOrDefaultInt("graphql.port", 8080)
}

type Config struct {
	App     AppConfig
	Redis   RedisConfig
	GraphQL GraphQlConfig
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
	config.App.parse()

	return config
}
