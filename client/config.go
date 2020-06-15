package client

import (
	"github.com/spf13/viper"
)

// Config is the logger configuration struct.
type config struct {
	address string
	crtPath string
}

// newConfig returns the configuration object.
func newConfig() *config {
	config := &config{
		address: viper.GetString("server_address"),
		crtPath: viper.GetString("server_crt_path"),
	}
	return config
}
