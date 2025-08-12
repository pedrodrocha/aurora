// Package config manages Aurora's configuration logic
package config

import (
	"fmt"

	"github.com/joho/godotenv"
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

func Load() (*Config, error) {
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}

func Generate() {
	fmt.Println("...generating config")
}

func Import() {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found
		} else {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}
}

func Exists() bool {
	config := viper.ConfigFileUsed()

	return config != ""
}

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./.aurora/")

	Import()

	godotenv.Load()
	viper.AutomaticEnv()

	viper.SetDefault("provider.type", "postgres")
	viper.SetDefault("provider.postgres.port", 5432)
	viper.SetDefault("provider.postgres.schema", "public")

	viper.BindEnv("provider.type", "PROVIDER_PROVIDER")

	viper.BindEnv("provider.postgres.host", "PROVIDER_POSTGRES_HOST")
	viper.BindEnv("provider.postgres.port", "PROVIDER_POSTGRES_PORT")
	viper.BindEnv("provider.postgres.user", "PROVIDER_POSTGRES_USER")
	viper.BindEnv("provider.postgres.host", "PROVIDER_POSTGRES_PASSWORD")
	viper.BindEnv("provider.postgres.host", "PROVIDER_POSTGRES_DATABASE")
	viper.BindEnv("provider.postgres.host", "PROVIDER_POSTGRES_SCHEMA")

	fmt.Println("Provider:", viper.GetString("provider.postgres.host"))
}
