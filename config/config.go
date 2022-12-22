package config

import (
	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath("./config/")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
}

func GetConfigs() error {
	err := viper.ReadInConfig()
	return err
}
