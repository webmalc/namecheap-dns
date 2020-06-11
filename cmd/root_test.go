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
	cr := NewCommandRouter(m, r)
	os.Args = []string{"invalid", "invalid"}
	m.On("Error", mock.Anything).Return(nil).Once()
	cr.Run()
	m.AssertExpectations(t)
}

// Should create a command router object.
func TestNewCommandRouter(t *testing.T) {
	m := &mocks.ErrorLogger{}
	r := &mocks.Changer{}
	cr := NewCommandRouter(m, r)
	assert.Equal(t, m, cr.logger)
	assert.Equal(t, r, cr.changer)
	assert.NotNil(t, cr.rootCmd)
}

func TestCommandRouter_change(t *testing.T) {
	r := &mocks.Changer{}
	cr := NewCommandRouter(&mocks.ErrorLogger{}, r)
	args := []string{"127.0.0.1"}
	r.On("Change", args[0]).Return(nil).Once()
	cr.change(&cobra.Command{}, args)
	r.AssertExpectations(t)
}

// Setups the tests.
func TestMain(m *testing.M) {
	test.Run(m)
}
