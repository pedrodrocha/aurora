// Package config provide a interactive aurora configuration setup
package config

import (
	"github.com/charmbracelet/huh"
	"github.com/pedro/aurora/internal/config"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	configCommand := &cobra.Command{
		Use:   "config",
		Short: "Generate configuration file for Aurora",
		Long:  "Generate interactively a config file for Aurora",
		Run: func(cmd *cobra.Command, args []string) {
			var confirm bool
			exists := config.Exists()

			if exists {
				huh.NewConfirm().
					Title("⚠️  Existing configuration detected").
					Description("Any changes will overwrite your current Aurora configuration. Are you sure you want to continue?").
					Affirmative("Continue").
					Negative("Cancel").
					Value(&confirm).
					WithTheme(huh.ThemeBase16()).
					Run()
			}

			if !exists || confirm {
				config.ExecuteForm()
			}

			// config.Generate()
		},
	}

	return configCommand
}
