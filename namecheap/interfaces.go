package namecheap

import "net"

// Logger logs the information
type logger interface {
	Infof(format string, args ...interface{})
	Error(args ...interface{})
}

// httpResponse is a http response
type httpResponse interface {
	ToXML(v interface{}) error
}

// httpClient does requests
type httpClient interface {
	Get(url string, v ...interface{}) (httpResponse, error)
}

// getter gets the DNS records
type getter interface {
	Get() (*apiResponse, error)
}

// setter sets the DNS records
type settter interface {
	Set(records *apiResponse, ip net.IP) error
}
