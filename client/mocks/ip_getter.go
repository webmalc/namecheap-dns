package mocks

import (
	"net"

	"github.com/stretchr/testify/mock"
)

// IPGetter logs errors.
type IPGetter struct {
	mock.Mock
}

// GetIP is a method mock
func (m *IPGetter) GetIP() (net.IP, error) {
	arg := m.Called()
	return arg.Get(0).(net.IP), nil
}
