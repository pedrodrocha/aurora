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

var defaults = map[string]any{
	"provider.type":            "postgres",
	"provider.postgres.port":   5432,
	"provider.postgres.schema": "public",
}

// Load reads and unmarshals the configuration into a Config struct.
// Returns an error if the configuration cannot be unmarshaled.
func Load() (*Config, error) {
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}

// Generate creates a new configuration file with default values.
// Currently prints a placeholder message - implementation pending.
func Generate() {
	fmt.Println("...generating config")
}

// Exists checks if a configuration file exists and has been loaded.
// Returns true if a config file path is set in viper.
func Exists() bool {
	config := viper.ConfigFileUsed()
	return config != ""
}

// Init initializes the configuration system.
// Returns an error if configuration fails to load.
func Init() error {
	if err := setupViper(); err != nil {
		return fmt.Errorf("viper setup failed: %w", err)
	}

	if err := loadConfig(); err != nil {
		return fmt.Errorf("config load failed: %w", err)
	}

	if err := loadEnv(); err != nil {
		return fmt.Errorf("env load failed: %w", err)
	}

	setDefaults()
	return nil
}

func setupViper() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./.aurora/")
	return nil
}

func loadConfig() error {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return fmt.Errorf("config file error: %w", err)
		}
		// Config file not found is not an error
	}
	return nil
}

func loadEnv() error {
	if err := godotenv.Load(); err != nil {
		log.Printf("warning: .env file not found: %v", err)
	}

	viper.AutomaticEnv()

	if err := bindEnvVars(); err != nil {
		return fmt.Errorf("env binding failed: %w", err)
	}

	return resolveEnvVars()
}

func bindEnvVars() error {
	for key, env := range envBindings {
		if err := viper.BindEnv(key, env); err != nil {
			return fmt.Errorf("failed to bind %s to %s: %w", key, env, err)
		}
	}
	return nil
}

func resolveEnvVars() error {
	all := viper.AllKeys()

	for _, key := range all {
		val := viper.GetString(key)

		if envName, ok := strings.CutPrefix(val, "ENV::"); ok {
			if envVal, ok := os.LookupEnv(envName); ok {
				viper.Set(key, envVal)
			} else {
				return fmt.Errorf("required env var %s not set (referenced by %s)", envName, key)
			}
		}
	}
	return nil
}

func setDefaults() {
	for key, value := range defaults {
		viper.SetDefault(key, value)
	}
}
