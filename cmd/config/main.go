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
			var (
				selectedProvider provider.Provider
				host             string
				port             string
				user             string
				password         string
			)

			providerOptions := []huh.Option[provider.Provider]{}
			for _, p := range provider.All() {
				providerOptions = append(providerOptions, huh.NewOption(p.Label(), p))
			}

			form := huh.NewForm(
				huh.NewGroup(
					huh.NewSelect[provider.Provider]().
						Title("Select a Provider").
						Options(providerOptions...),

					huh.NewInput().
						Title("Host").
						Value(&host),

					huh.NewInput().
						Title("Port").
						Value(&port),
					huh.NewInput().
						Title("User").
						Value(&user),
					huh.NewInput().
						Title("Password").
						Value(&password),
				),
			)

			if err := form.Run(); err != nil {
				fmt.Println("Error running form:", err)
				return
			}

			fmt.Println("Provider: ", selectedProvider)
			fmt.Println("Host: ", host)
			fmt.Println("Port: ", port)
			fmt.Println("User: ", user)
			fmt.Println("Password: ", password)
		},
	}

	return configCommand
}
