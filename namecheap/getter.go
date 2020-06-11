package namecheap

import (
	"github.com/pkg/errors"
)

const apiGetMethod = "namecheap.domains.dns.getHosts"
const statusOK = "OK"

// apiGetter gets the DNS records through the API
type apiGetter struct {
	config *Config
	http   httpClient
}

// Get gets the DNS records.
func (a *apiGetter) Get() (*apiResponse, error) {
	r, err := a.http.Get(a.config.BaseURL + apiGetMethod)
	if err != nil {
		return nil, errors.Wrap(err, "api getter")
	}
	var apiResponse apiResponse
	err = r.ToXML(&apiResponse)
	if err != nil {
		return nil, errors.Wrap(err, "api getter")
	}
	if apiResponse.Status != statusOK {
		return nil, errors.Wrap(errors.New("invalid response"), "api getter")
	}

	return &apiResponse, nil
}

// newApiGetter creates a new api getter object
func newAPIGetter(c *Config, http httpClient) *apiGetter {
	return &apiGetter{config: c, http: http}
}
