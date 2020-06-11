package namecheap

import (
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
