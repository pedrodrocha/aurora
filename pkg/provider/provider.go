// Package provider manages the supported database providers for Aurora.
package provider

import "fmt"

type Provider string

const (
	Unknown  Provider = "unknown"
	Postgres Provider = "postgres"
)

var providerLabels = map[Provider]string{
	Postgres: "Postgres",
}

func (p Provider) Label() string {
	if label, ok := providerLabels[p]; ok {
		return label
	}
	return "Unknown"
}

func IsSupported(p Provider) bool {
	return p != Unknown
}

func Parse(s string) (Provider, error) {
	switch Provider(s) {
	case Postgres:
		return Postgres, nil
	default:
		return Unknown, fmt.Errorf("unknown provider %q", s)
	}
}

func (p *Provider) UnmarshalText(text []byte) error {
	val, err := Parse(string(text))
	if err != nil {
		return err
	}
	*p = val
	return nil
}

func All() []Provider {
	return []Provider{Postgres}
}
