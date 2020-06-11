package namecheap

import (
	"fmt"
	"net"

	"github.com/pkg/errors"
)

// Changer change the DNS
type Changer struct {
	logger logger
	config *Config
	getter getter
}

// getIP gets IP address from the string
func (c *Changer) getIP(address string) (net.IP, error) {
	ip := net.ParseIP(address)
	if ip == nil {
		return nil, errors.Wrap(errors.New("invalid IP address"), "changer")
	}
	return ip, nil
}

// Change change the DNS
func (c *Changer) Change(address string) {
	ip, err := c.getIP(address)
	if err != nil {
		c.logger.Error(err)
		return
	}
	records, err := c.getter.Get()
	if err != nil {
		c.logger.Error(err)
		return
	}
	c.logger.Infof("Got the records: %v", records)
	fmt.Println(records)
	fmt.Println(ip)
}

// NewChanger creates a new Changer
func NewChanger(l logger) *Changer {
	c := newConfig()
	http := newHTTPClient()
	return &Changer{
		logger: l, config: c, getter: newAPIGetter(c, http),
	}
}
