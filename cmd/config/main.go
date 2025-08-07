// Package config provide a interactive aurora configuration setup
package config

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/pedro/aurora/pkg/provider"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	configCommand := &cobra.Command{
		Use:   "config",
		Short: "Generate configuration file for Aurora",
		Long:  "Generate interactively a config file for Aurora",
		Run: func(cmd *cobra.Command, args []string) {
			var selectedProvider provider.Provider

			providerOptions := []huh.Option[provider.Provider]{}
			for _, p := range provider.All() {
				providerOptions = append(providerOptions, huh.NewOption(p.Label(), p))
			}

			form := huh.NewForm(
				huh.NewGroup(
					huh.NewSelect[provider.Provider]().
						Title("Select a Provider").
						Options(providerOptions...),
				),
			)

			if err := form.Run(); err != nil {
				fmt.Println("Error running form:", err)
				return
			}

			fmt.Println("Selected Provider:", selectedProvider)
		},
	}

	return configCommand
}
