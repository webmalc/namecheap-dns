package namecheap

import (
	"net"

	"github.com/stretchr/testify/mock"
)

//golint:ignore

// getterErrorMock is a mock object.
type getterErrorMock struct {
	mock.Mock
}

// Get is the mock method
func (c *getterErrorMock) Get() (*apiResponse, error) {
	arg := c.Called()
	return arg.Get(0).(*apiResponse), arg.Get(1).(error)
}

// getterMock is a mock object.
type getterMock struct {
	mock.Mock
}

// Get is the mock method
func (c *getterMock) Get() (*apiResponse, error) {
	arg := c.Called()
	return arg.Get(0).(*apiResponse), nil
}

type setterErrorMock struct {
	mock.Mock
}

// Get is the mock method
func (c *setterErrorMock) Set(r *apiResponse, ip net.IP) error {
	arg := c.Called(r, ip)
	return arg.Get(0).(error)
}
