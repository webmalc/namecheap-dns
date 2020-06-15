package mocks

import (
	"github.com/stretchr/testify/mock"
)

// Changer mocks the object
type Changer struct {
	mock.Mock
}

// Change is method mock
func (r *Changer) Change(address string) {
	r.Called(address)
}
