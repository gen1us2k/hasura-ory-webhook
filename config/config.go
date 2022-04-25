package config

import "github.com/kelseyhightower/envconfig"

// HookConfig struct stores configuration
// for the webhook passed via environment variables
type HookConfig struct {
	OrySDKURL string `envconfig:"ORY_SDK_URL"`
}

// Parse parses envoriment variables and returns HookConfig structure
func Parse() (*HookConfig, error) {
	var h HookConfig
	err := envconfig.Process("", &h)
	return &h, err
}
