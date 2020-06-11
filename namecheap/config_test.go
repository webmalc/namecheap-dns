package namecheap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Should return the config object
func TestNewConfig(t *testing.T) {
	c := newConfig()
	assert.Contains(t, c.BaseURL, "namecheap.com/xml.response?ApiUser=user&")
}
