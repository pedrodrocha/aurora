// Package provider manages the supported database providers for Aurora.
package provider

import "fmt"

type Provider int

const Unknown Provider = -1

const (
	Postgres Provider = iota
)

var providerIds = map[Provider]string{
	Postgres: "postgres",
}

var providerValues = map[string]Provider{
	"postgres": Postgres,
}

var providerLabels = map[Provider]string{
	Postgres: "Postgres",
}

func (p Provider) String() string {
	if name, ok := providerIds[p]; ok {
		return name
	}
	return "unknown"
}

func (p Provider) Label() string {
	if label, ok := providerLabels[p]; ok {
		return label
	}
	return "Unknown"
}

func IsSupported(p Provider) bool {
	_, ok := providerIds[p]
	return ok
}

func Parse(s string) (Provider, error) {
	p, ok := providerValues[s] // s is string, keys are string
	if !ok {
		return Unknown, fmt.Errorf("unknown provider %q", s)
	}
	return p, nil
}

func All() []Provider {
	providers := make([]Provider, 0, len(providerIds))
	for p := range providerIds {
		providers = append(providers, p)
	}
	return providers
}
