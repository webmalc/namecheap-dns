package namecheap

import (
	"github.com/spf13/viper"
)

// Config is the logger configuration struct.
type Config struct {
	BaseURL string
}

// NewConfig returns the configuration object.
func newConfig() *Config {
	config := &Config{
		BaseURL: viper.GetString("base_url"),
	}
	return config
}
