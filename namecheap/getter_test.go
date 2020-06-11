package namecheap

import (
	"errors"
	"testing"

	"github.com/imroc/req"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Should return an error
func Test_apiGetter_Get_Invalid_Request(t *testing.T) {
	c := newConfig()
	h := &httpClientErrorMock{}
	g := newAPIGetter(c, h)
	h.On("Get", mock.Anything, mock.Anything).Return(
		&req.Resp{}, errors.New("test error"),
	).Once()
	_, err := g.Get()
	assert.Error(t, err)
	h.AssertExpectations(t)
}

// Should return an error
func Test_apiGetter_Get_Invalid_XMl(t *testing.T) {
	c := newConfig()
	h := &httpClientMock{}
	g := newAPIGetter(c, h)
	r := &httpResponseErrorMock{}
	h.On("Get", mock.Anything, mock.Anything).Return(r, nil).Once()
	r.On("ToXML", mock.Anything).Return(errors.New("test")).Once()
	_, err := g.Get()
	assert.Error(t, err)
	h.AssertExpectations(t)
	r.AssertExpectations(t)
}

// Should return an error
func Test_apiGetter_Get_Invalid_Status(t *testing.T) {
	c := newConfig()
	h := &httpClientMock{}
	g := newAPIGetter(c, h)
	r := &httpResponseMock{}
	h.On("Get", mock.Anything, mock.Anything).Return(r, nil).Once()
	r.On("ToXML", mock.Anything).Return("").Once()
	_, err := g.Get()
	assert.Error(t, err)
	h.AssertExpectations(t)
	r.AssertExpectations(t)
}

// Should create a new api getter
func Test_newAPIGetter(t *testing.T) {
	c := newConfig()
	h := newHTTPClient()
	g := newAPIGetter(c, h)
	assert.Equal(t, c, g.config)
	assert.Equal(t, h, g.http)
}
