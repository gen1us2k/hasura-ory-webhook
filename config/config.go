package config

import "github.com/kelseyhightower/envconfig"

const (
	// Development environment constant
	Development string = "development"
	// Production environment constant
	Production string = "production"
)

// HookConfig struct stores configuration
// for the webhook passed via environment variables
type HookConfig struct {
	OrySDKURL   string `envconfig:"ORY_SDK_URL"`
	Environment string `envconfig:"ENV" default:"development"`
}

// Parse parses envoriment variables and returns HookConfig structure
func Parse() (*HookConfig, error) {
	var h HookConfig
	err := envconfig.Process("", &h)
	return &h, err
}
