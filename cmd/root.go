// Package cmd provides the entrypoint for aurora cli
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	configCmd "github.com/pedro/aurora/cmd/config"
	"github.com/pedro/aurora/internal/config"
)

var rootCmd = &cobra.Command{
	Use:   "aurora",
	Short: "A system for bidirectional database schema translation",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(configCmd.New())

	config.Define()

	config.Import()
}
