package mocks

import (
	"github.com/stretchr/testify/mock"
)

// Server mocks the object
type Server struct {
	mock.Mock
}

// Run is method mock
func (m *Server) Run() {
	m.Called()
}
