package client

import "net"

// Logger logs the information
type logger interface {
	Infof(format string, args ...interface{})
	Error(args ...interface{})
}

type ipGetter interface {
	GetIP() (net.IP, error)
}
