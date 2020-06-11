package test

import (
	"os"
	"testing"

	"github.com/webmalc/namecheap-dns/common/config"
)

// Setups the tests
func setUp() {
	os.Setenv("NDNS_ENV", "test")
	config.Setup()
}

// Run setups, runs and teardowns the tests
func Run(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}
