package namecheap

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
