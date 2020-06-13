package namecheap

import (
	"errors"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/webmalc/namecheap-dns/common/test"
)

// Should get an IP address from the string
func TestChanger_getIP(t *testing.T) {
	l := &loggerMock{}
	c := NewChanger(l)
	_, err := c.getIP("invalid")
	assert.Error(t, err)

	ip, err := c.getIP("127.0.0.1")
	assert.Nil(t, err)
	assert.Equal(t, net.ParseIP("127.0.0.1"), ip)
}

// Should log the error
func TestChanger_Change_Invalid_IP(t *testing.T) {
	l := &loggerMock{}
	c := NewChanger(l)
	l.On("Error", mock.Anything).Return(nil).Once()
	c.Change("invalid")
	l.AssertExpectations(t)
}

// Should log the error
func TestChanger_Change_Invalid_Getter(t *testing.T) {
	l := &loggerMock{}
	c := NewChanger(l)
	g := &getterErrorMock{}
	c.getter = g
	l.On("Error", mock.Anything).Return(nil).Once()
	g.On("Get").Return(&apiResponse{}, errors.New("test")).Once()
	c.Change("127.0.0.1")
	l.AssertExpectations(t)
	g.AssertExpectations(t)
}

// Should log the error
func TestChanger_Change_Invalid_Setter(t *testing.T) {
	l := &loggerMock{}
	c := NewChanger(l)
	g := &getterMock{}
	s := &setterErrorMock{}
	c.getter = g
	c.setter = s
	ip := "1.1.1.1"
	resp := &apiResponse{Status: "OK"}
	g.On("Get").Return(resp).Once()
	l.On("Infof", mock.Anything, mock.Anything).Return(nil).Once()
	s.On("Set", resp, net.ParseIP(ip)).Return(errors.New("test")).Once()
	l.On("Error", mock.Anything).Return(nil).Once()
	c.Change(ip)
	l.AssertExpectations(t)
	g.AssertExpectations(t)
	s.AssertExpectations(t)
}

// Should create a new changer object.
func TestNewChanger(t *testing.T) {
	l := &loggerMock{}
	c := NewChanger(l)
	assert.Equal(t, l, c.logger)
}

// Setups the tests.
func TestMain(m *testing.M) {
	test.Run(m)
}
