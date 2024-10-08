package config

import (
	"github.com/spf13/viper"
	"time"
)
import log "github.com/sirupsen/logrus"

func viperGetOrDefaultBool(
	key string,
	defaultValue bool,
) bool {
	viper.SetDefault(key, defaultValue)
	return viper.GetBool(key)
}

func viperGetOrDefaultString(
	key string,
	defaultValue string,
) string {
	viper.SetDefault(key, defaultValue)
	return viper.GetString(key)
}

func viperGetOrDefaultInt(
	key string,
	defaultValue int,
) int64 {
	viper.SetDefault(key, defaultValue)
	return viper.GetInt64(key)
}

func viperGetOrDefaultUint(
	key string,
	defaultValue uint64,
) uint64 {
	viper.SetDefault(key, defaultValue)
	return viper.GetUint64(key)
}

func viperGetOrDefaultTimeDuration(
	key string,
	defaultValue string,
) time.Duration {
	viper.SetDefault(key, defaultValue)
	d, err := time.ParseDuration(viper.GetString(key))
	if err != nil {
		log.Fatalf("provided value '%s' cannot be transformed to [time.Duration]", viper.GetString(key))
	}
	return d
}

func viperGetOrDefaultFloat(key string, defaultValue float64) float64 {
	viper.SetDefault(key, defaultValue)
	return viper.GetFloat64(key)
}
