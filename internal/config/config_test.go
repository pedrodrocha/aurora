package config

import (
	"testing"

	"github.com/pedro/aurora/pkg/provider"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func resetViper() {
	viper.Reset()
	setDefaults()
}

func TestLoad(t *testing.T) {
	t.Run("successful load", func(t *testing.T) {
		resetViper()
		viper.Set("provider.provider", "postgres")
		viper.Set("provider.postgres.host", "localhost")

		cfg, err := Load()
		require.NoError(t, err)
		assert.Equal(t, provider.Postgres, cfg.Provider.Type)
		assert.Equal(t, "localhost", cfg.Provider.Postgres.Host)
	})

	t.Run("invalid config", func(t *testing.T) {
		resetViper()
		viper.Set("provider.provider", "invalid_provider")

		_, err := Load()
		require.Error(t, err)
		assert.Contains(t, err.Error(), "unknown provider")
	})
}

func TestExists(t *testing.T) {
	t.Run("config exists", func(t *testing.T) {
		viper.SetConfigFile("test.toml")
		defer viper.Reset()

		assert.True(t, Exists())
	})

	t.Run("config missing", func(t *testing.T) {
		viper.Reset()
		assert.False(t, Exists())
	})
}

func TestInit(t *testing.T) {
	t.Run("successful init", func(t *testing.T) {
		viper.Reset()

		err := Init()
		require.NoError(t, err)
		assert.Equal(t, "postgres", viper.GetString("provider.provider"))
	})
}

func TestBindEnvVars(t *testing.T) {
	t.Run("successful binding", func(t *testing.T) {
		err := bindEnvVars()
		require.NoError(t, err)
	})
}

func TestResolveEnvVars(t *testing.T) {
	t.Run("successful resolution", func(t *testing.T) {
		t.Setenv("TEST_VAR", "test_value")
		viper.Set("test.key", "ENV::TEST_VAR")

		err := resolveEnvVars()
		require.NoError(t, err)
		assert.Equal(t, "test_value", viper.GetString("test.key"))
	})

	t.Run("missing env var", func(t *testing.T) {
		viper.Set("test.key", "ENV::MISSING_VAR")

		err := resolveEnvVars()
		require.Error(t, err)
		assert.Contains(t, err.Error(), "required env var")
	})
}

func TestSetDefaults(t *testing.T) {
	viper.Reset()
	setDefaults()

	assert.Equal(t, "postgres", viper.GetString("provider.provider"))
	assert.Equal(t, 5432, viper.GetInt("provider.postgres.port"))
	assert.Equal(t, "public", viper.GetString("provider.postgres.schema"))
}
