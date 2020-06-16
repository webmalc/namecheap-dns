package client

import (
	"net"
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
	i := &mocks.IPGetter{}
	ip := "3.3.2.2"
	i.On("GetIP").Return(net.ParseIP(ip)).Once()
	c.On("Change", ip).Return(nil).Once()
	l.On("Infof", mock.Anything, mock.Anything).Return(nil).Twice()
	go s.Run()
	client := NewClient(l, i)
	client.Request()
	i.AssertExpectations(t)
	c.AssertExpectations(t)
	l.AssertExpectations(t)
}

// Should return a new client
func TestNewClient(t *testing.T) {
	l := &mocks.Logger{}
	i := &mocks.IPGetter{}
	c := NewClient(l, i)
	assert.Equal(t, c.logger, l)
	assert.Equal(t, c.ipGetter, i)
}
