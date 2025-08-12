// Package config manages Aurora's configuration logic
package config

import (
	"fmt"
	"log"
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

var envBindings = map[string]string{
	"provider.type":              "PROVIDER_TYPE",
	"provider.postgres.host":     "PROVIDER_POSTGRES_HOST",
	"provider.postgres.port":     "PROVIDER_POSTGRES_PORT",
	"provider.postgres.user":     "PROVIDER_POSTGRES_USER",
	"provider.postgres.password": "PROVIDER_POSTGRES_PASSWORD",
	"provider.postgres.database": "PROVIDER_POSTGRES_DATABASE",
	"provider.postgres.schema":   "PROVIDER_POSTGRES_SCHEMA",
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

	bindEnvVars()
	resolveEnvVars()
}

func bindEnvVars() {
	for key, env := range envBindings {
		if err := viper.BindEnv(key, env); err != nil {
			log.Fatalf("bind env error: %v", err)
		}
	}
}

func resolveEnvVars() {
	all := viper.AllKeys()

	for _, key := range all {
		val := viper.GetString(key)

		if envName, ok := strings.CutPrefix(val, "ENV::"); ok {
			if envVal, ok := os.LookupEnv(envName); ok {
				viper.Set(key, envVal)
			}
		}
	}
}

func setDefaults() {
	viper.SetDefault("provider.type", "postgres")
	viper.SetDefault("provider.postgres.port", 5432)
	viper.SetDefault("provider.postgres.schema", "public")
}
