package mocks

import (
	"github.com/stretchr/testify/mock"
)

// Client mocks the object
type Client struct {
	mock.Mock
}

// Request is method mock
func (m *Client) Request(ip string) {
	m.Called(ip)
}
