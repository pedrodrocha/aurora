// Package config manages Aurora's configuration logic
package config

import (
	"errors"
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/pedro/aurora/pkg/provider"
)

func ExecuteForm() {
	var (
		selectedProvider provider.Provider
		host             string
		port             string
		user             string
		password         string
		database         string
		schema           = "public"
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
				Title("Database").
				Value(&database).
				Validate(func(str string) error {
					if str == "" {
						return errors.New("database is required")
					}

					return nil
				}),

			huh.NewInput().
				Title("Schema").
				Value(&schema),

			huh.NewInput().
				Title("Host").
				Value(&host).
				Validate(func(str string) error {
					if str == "" {
						return errors.New("host is required")
					}

					return nil
				}),

			huh.NewInput().
				Title("Port").
				Value(&port).
				Validate(func(str string) error {
					if str == "" {
						return errors.New("port is required")
					}
					return nil
				}),
			huh.NewInput().
				Title("User").
				Value(&user).
				Validate(func(str string) error {
					if str == "" {
						return errors.New("user is required")
					}
					return nil
				}),

			huh.NewInput().
				Title("Password").
				Value(&password),
		),
	).WithTheme(huh.ThemeBase16())

	if err := form.Run(); err != nil {
		fmt.Println("Error running form:", err)
		return
	}

	fmt.Println("Provider: ", selectedProvider)
	fmt.Println("Host: ", host)
	fmt.Println("Port: ", port)
	fmt.Println("User: ", user)
	fmt.Println("Password: ", password)
	fmt.Println("Database: ", database)
	fmt.Println("Schema: ", schema)
}
