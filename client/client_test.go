package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/webmalc/namecheap-dns/client/mocks"
	"github.com/webmalc/namecheap-dns/server"
)

func TestClient_Request(t *testing.T) {
	l := &mocks.Logger{}
	c := &mocks.Changer{}
	s := server.NewServer(l, c)
	ip := "3.3.2.2"
	c.On("Change", ip).Return(nil).Once()
	l.On("Infof", mock.Anything, mock.Anything).Return(nil).Once()
	go s.Run()
	client := NewClient(l)
	client.Request(ip)
	c.AssertExpectations(t)
	l.AssertExpectations(t)
}

// Should return a new client
func TestNewClient(t *testing.T) {
	l := &mocks.Logger{}
	c := NewClient(l)
	assert.Equal(t, c.logger, l)
}
