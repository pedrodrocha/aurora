// Package config provide a interactive aurora configuration setup
package config

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	configCommand := &cobra.Command{
		Use:   "config",
		Short: "Generate configuration file for Aurora",
		Long:  "Generate interactively a config file for Aurora",
		Run: func(cmd *cobra.Command, args []string) {
			var name string

			huh.NewInput().
				Title("What’s your name?").
				Value(&name).
				Run() // this is blocking...

			fmt.Printf("Hey, %s!\n", name)
		},
	}

	return configCommand
}
