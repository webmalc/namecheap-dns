package server

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	pb "github.com/webmalc/namecheap-dns/protos/changer"
	"github.com/webmalc/namecheap-dns/server/mocks"
)

func TestServer_Change(t *testing.T) {
	l := &mocks.Logger{}
	c := &mocks.Changer{}
	s := NewServer(l, c)
	ip := "2.2.2.2"
	c.On("Change", ip).Return(nil).Once()
	_, err := s.Change(context.TODO(), &pb.ChangeRequest{Ip: ip})
	assert.Nil(t, err)
	c.AssertExpectations(t)
}

func TestServer_Run(t *testing.T) {
	l := &mocks.Logger{}
	c := &mocks.Changer{}
	s := NewServer(l, c)
	go s.Run()
}

func TestNewServer(t *testing.T) {
	l := &mocks.Logger{}
	c := &mocks.Changer{}
	s := NewServer(l, c)
	assert.Equal(t, s.logger, l)
	assert.Equal(t, s.changer, c)
}
