package ip

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/namecheap-dns/ip/mocks"
)

// Should return a public IP
func TestGetter_GetIP(t *testing.T) {
	l := &mocks.Logger{}
	i := NewGetter(l)
	ip, err := i.GetIP()
	assert.NotNil(t, ip)
	assert.Nil(t, err)
}

// Should create a public IP getter instance
func TestNewGetter(t *testing.T) {
	l := &mocks.Logger{}
	i := NewGetter(l)
	assert.Equal(t, i.logger, l)
}
