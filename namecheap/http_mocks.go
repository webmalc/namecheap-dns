package namecheap

import (
	"github.com/stretchr/testify/mock"
)

// httpClientErrorMock is a mock object.
type httpClientErrorMock struct {
	mock.Mock
}

// Get is the mock method
func (c *httpClientErrorMock) Get(
	url string, v ...interface{},
) (httpResponse, error) {
	arg := c.Called(url, v)
	return arg.Get(0).(httpResponse), arg.Get(1).(error)
}

type httpClientMock struct {
	mock.Mock
}

// Get is the mock method
func (c *httpClientMock) Get(
	url string, v ...interface{},
) (httpResponse, error) {
	arg := c.Called(url, v)
	return arg.Get(0).(httpResponse), nil
}

// httpResponseErrorMock is a mock object.
type httpResponseErrorMock struct {
	mock.Mock
}

// ToXML is the mock method
func (r *httpResponseErrorMock) ToXML(v interface{}) error {
	arg := r.Called(v)
	return arg.Get(0).(error)
}

// httpResponseMock is a mock object.
type httpResponseMock struct {
	mock.Mock
}

// ToXML is the mock method
func (r *httpResponseMock) ToXML(v interface{}) error {
	r.Called(v)
	return nil
}
