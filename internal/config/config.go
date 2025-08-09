// Package config manages Aurora's configuration logic
package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Generate() {
	fmt.Println("...generating config")
}

func Import() {
	viper.ReadInConfig()
}

func Exists() bool {
	config := viper.ConfigFileUsed()

	return config != ""
}

func Define() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./.aurora/")
}
