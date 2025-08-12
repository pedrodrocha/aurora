// Package config manages Aurora's configuration logic
package config

import (
	"fmt"
	"os"
	"strings"

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

func Exists() bool {
	config := viper.ConfigFileUsed()

	return config != ""
}

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./.aurora/")

	read()
	mergeEnv()
	setDefaults()
}

func read() {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found
		} else {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}
}

func mergeEnv() {
	godotenv.Load()
	viper.AutomaticEnv()

	viper.BindEnv("provider.type", "PROVIDER_PROVIDER")

	viper.BindEnv("provider.postgres.host", "PROVIDER_POSTGRES_HOST")
	viper.BindEnv("provider.postgres.port", "PROVIDER_POSTGRES_PORT")
	viper.BindEnv("provider.postgres.user", "PROVIDER_POSTGRES_USER")
	viper.BindEnv("provider.postgres.host", "PROVIDER_POSTGRES_PASSWORD")
	viper.BindEnv("provider.postgres.host", "PROVIDER_POSTGRES_DATABASE")
	viper.BindEnv("provider.postgres.host", "PROVIDER_POSTGRES_SCHEMA")

	resolveEnvVars()
}

func setDefaults() {
	viper.SetDefault("provider.type", "postgres")
	viper.SetDefault("provider.postgres.port", 5432)
	viper.SetDefault("provider.postgres.schema", "public")
}

func resolveEnvVars() {
	// iterate over all keys and for string values that start with "ENV::" replace them
	for _, key := range viper.AllKeys() {
		val := viper.Get(key)
		strVal, ok := val.(string)
		if !ok {
			continue
		}
		if strings.HasPrefix(strVal, "ENV::") {
			envName := strings.TrimPrefix(strVal, "ENV::")
			if envVal, ok := os.LookupEnv(envName); ok {
				viper.Set(key, envVal)
			}
		}
	}
}
