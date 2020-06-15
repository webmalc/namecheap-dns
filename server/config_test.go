package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/namecheap-dns/common/test"
)

// Should return the config object
func TestNewConfig(t *testing.T) {
	c := newConfig()
	assert.Contains(t, c.address, ":50052")
	assert.Contains(t, c.keyPath, "server.key")
	assert.Contains(t, c.crtPath, "server.crt")
}

// Setups the tests.
func TestMain(m *testing.M) {
	test.Run(m)
}
