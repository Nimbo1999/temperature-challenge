package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadEnvVariables() error {
	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../../")
	viper.AddConfigPath("..")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return fmt.Errorf("could not find the configuration file:\n%v", err)
		} else {
			return err
		}
	}

	return nil
}
