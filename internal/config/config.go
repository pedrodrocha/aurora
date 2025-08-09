// Package config manages Aurora's configuration logic
package config

import (
	"fmt"

	"github.com/pedro/aurora/pkg/provider"
	"github.com/spf13/viper"
)

type Config struct {
	Provider struct {
		Type     provider.Provider `mapstructure:"provider"`
		Postgres PostgresConfig    `mapstructure:"postgres"`
	} `mapstructure:"provider"`
}

type PostgresConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	Schema   string `mapstructure:"schema"`
}

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

func Load() (*Config, error) {
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./.aurora/")

	viper.ReadInConfig()

	viper.SetDefault("provider.type", "postgres")
	viper.SetDefault("provider.postgres.port", 5432)
	viper.SetDefault("provider.postgres.schema", "public")
}
