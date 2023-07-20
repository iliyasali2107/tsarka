package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	ServerUrl        string        `mapstructure:"SERVER_URL"`
	RedisUrl         string        `mapstructure:"REDIS_URL"`
	HashMaxRequests  int           `mapstructure:"HASH_MAX_REQUESTS"`
	HashInternval    time.Duration `mapstructure:"HASH_INTERVAL"`
	HashCalcDuration time.Duration `mapstructure:"HASH_CALC_DURATION"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./internal/config/envs")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	viper.SetDefault("SERVER_URL", ":50052")
	viper.SetDefault("REDIS_URL", ":6379")
	viper.SetDefault("HASH_MAX_REQUESTS", 5)
	viper.SetDefault("HASH_INTERVAL", 5*time.Second)
	viper.SetDefault("HASH_CALC_DURATION", 1*time.Minute)

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	fmt.Println(config)
	return
}
