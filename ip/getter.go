package ip

import (
	"net"

	"github.com/imroc/req"
	"github.com/pkg/errors"
)

const serviceURL = "https://ipecho.net/plain"

// Getter is the public IP getter
type Getter struct {
	logger logger
}

// GetIP gets the public IP
func (g *Getter) GetIP() (net.IP, error) {
	resp, err := req.Get(serviceURL)
	if err != nil {
		g.logger.Error(errors.Wrap(err, "ip getter"))
		return nil, err
	}
	ip, err := resp.ToString()
	if err != nil {
		g.logger.Error(errors.Wrap(err, "ip getter"))
		return nil, err
	}
	return net.ParseIP(ip), nil
}

// NewGetter creates a new getter instance
func NewGetter(l logger) *Getter {
	return &Getter{logger: l}
}
