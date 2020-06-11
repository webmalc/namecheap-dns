package namecheap

import (
	"github.com/imroc/req"
)

type client struct{}

func (c *client) Get(url string, v ...interface{}) (httpResponse, error) {
	return req.Get(url, v)
}

// newHttpClient creates a new http client
func newHTTPClient() httpClient {
	return &client{}
}
