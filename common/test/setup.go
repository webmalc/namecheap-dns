package test

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/webmalc/namecheap-dns/common/config"
)

// Setups the tests
func setUp() {
	os.Setenv("NDNS_ENV", "test")
	config.Setup()
	keyDir := viper.GetString("base_dir") + "common/config/cert_test/"
	viper.SetDefault("server_key_path", keyDir+"server.key")
	viper.SetDefault("server_crt_path", keyDir+"server.crt")
}

// Run setups, runs and teardowns the tests
func Run(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}
