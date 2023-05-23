package configs

import (
	"electro-item-management/domains"
	"fmt"

	"github.com/spf13/viper"
)

var (
	AppConfig domains.Config
)

func LoadConfig() *domains.Config {
	viper.SetConfigFile("env")

	// uncomment dev and comment local if using docker
	viper.SetConfigName("local")
	// viper.SetConfigName("dev")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err := viper.Unmarshal(&AppConfig)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return &AppConfig
}
