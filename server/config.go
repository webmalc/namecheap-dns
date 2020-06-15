package server

import (
	"github.com/spf13/viper"
)

// Config is the logger configuration struct.
type config struct {
	address string
	keyPath string
	crtPath string
}

func newConfig() *config {
	config := &config{
		address: viper.GetString("server_address"),
		keyPath: viper.GetString("server_key_path"),
		crtPath: viper.GetString("server_crt_path"),
	}
	return config
}
