package namecheap

import (
	"github.com/stretchr/testify/mock"
)

// loggerMock logs errors.
type loggerMock struct {
	mock.Mock
}

// Error is a method mock
func (m *loggerMock) Error(args ...interface{}) {
	m.Called(args...)
}

// Infof is a method mock
func (m *loggerMock) Infof(format string, args ...interface{}) {
	m.Called(format, args)
}
