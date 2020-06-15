package cmd

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/webmalc/namecheap-dns/cmd/mocks"
	"github.com/webmalc/namecheap-dns/common/test"
)

// Should run the root command and log an error.
func TestCommandRouter_Run(t *testing.T) {
	m := &mocks.ErrorLogger{}
	r := &mocks.Changer{}
	s := &mocks.Server{}
	c := &mocks.Client{}
	cr := NewCommandRouter(m, r, s, c)
	os.Args = []string{"invalid", "invalid"}
	m.On("Error", mock.Anything).Return(nil).Once()
	cr.Run()
	m.AssertExpectations(t)
}

// Should create a command router object.
func TestNewCommandRouter(t *testing.T) {
	m := &mocks.ErrorLogger{}
	r := &mocks.Changer{}
	s := &mocks.Server{}
	c := &mocks.Client{}
	cr := NewCommandRouter(m, r, s, c)
	assert.Equal(t, m, cr.logger)
	assert.Equal(t, r, cr.changer)
	assert.NotNil(t, cr.rootCmd)
}

// Should run the changer
func TestCommandRouter_change(t *testing.T) {
	m := &mocks.ErrorLogger{}
	r := &mocks.Changer{}
	s := &mocks.Server{}
	c := &mocks.Client{}
	cr := NewCommandRouter(m, r, s, c)
	args := []string{"127.0.0.1"}
	r.On("Change", args[0]).Return(nil).Once()
	cr.change(&cobra.Command{}, args)
	r.AssertExpectations(t)
}

// Should run the server
func TestCommandRouter_server(t *testing.T) {
	m := &mocks.ErrorLogger{}
	r := &mocks.Changer{}
	s := &mocks.Server{}
	c := &mocks.Client{}
	cr := NewCommandRouter(m, r, s, c)
	s.On("Run").Return(nil).Once()
	cr.runServer(&cobra.Command{}, []string{})
	s.AssertExpectations(t)
}

// Should run the request
func TestCommandRouter_request(t *testing.T) {
	m := &mocks.ErrorLogger{}
	r := &mocks.Changer{}
	s := &mocks.Server{}
	c := &mocks.Client{}
	cr := NewCommandRouter(m, r, s, c)
	args := []string{"1.1.1.1"}
	c.On("Request", args[0]).Return(nil).Once()
	cr.request(&cobra.Command{}, args)
	c.AssertExpectations(t)
}

// Setups the tests.
func TestMain(m *testing.M) {
	test.Run(m)
}
