package namecheap

import "fmt"

// Changer change the DNS
type Changer struct{}

// Change change the DNS
func (r *Changer) Change(address string) {
	fmt.Println(address)
}

// NewChanger creates a new Changer
func NewChanger() *Changer {
	return &Changer{}
}
