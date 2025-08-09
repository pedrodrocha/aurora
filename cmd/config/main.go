// Package config provide a interactive aurora configuration setup
package config

import (
	"fmt"

	"github.com/pedro/aurora/internal/config"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	configCommand := &cobra.Command{
		Use:   "config",
		Short: "Generate configuration file for Aurora",
		Long:  "Generate interactively a config file for Aurora",
		Run: func(cmd *cobra.Command, args []string) {
			exists := config.Exists()
			loaded, _ := config.Load()
			fmt.Println("exists: ", exists)
			fmt.Println("loaded", loaded)
			fmt.Println(loaded.Provider.Type.Label())

			// config.ExecuteForm()

			// config.Generate()
		},
	}

	return configCommand
}
