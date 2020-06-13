package namecheap

import (
	"fmt"
	"net"
	"net/url"

	"github.com/pkg/errors"
)

const apiSetMethod = "namecheap.domains.dns.setHosts"

// apiSetter sets the DNS records through the API
type apiSetter struct {
	config *Config
	http   httpClient
}

// update updates the records
func (a *apiSetter) updateRecords(records *apiResponse, ip net.IP) bool {
	updated := false
	addr := ip.String()
	for i, host := range records.Hosts {
		if host.Type == "A" && host.Address != addr {
			records.Hosts[i].Address = addr
			updated = true
		}
	}
	return updated
}

// generateURL generates the URL to updates the DNS records
func (a *apiSetter) generateURL(records *apiResponse) string {
	u := ""
	for i, host := range records.Hosts {
		i++
		u += fmt.Sprintf("HostName%d=%s&", i, host.Name)
		u += fmt.Sprintf("RecordType%d=%s&", i, host.Type)
		u += fmt.Sprintf("Address%d=%s&", i, host.Address)
		u += fmt.Sprintf("TTL%d=%d&", i, host.TTL)
		if host.Type == "MX" {
			u += fmt.Sprintf("MXPref%d=%d&", i, host.MXPref)
		}
	}
	return a.config.BaseURL + apiSetMethod + "&" + url.PathEscape(u)
}

// Get gets the DNS records.
func (a *apiSetter) Set(records *apiResponse, ip net.IP) error {
	updated := a.updateRecords(records, ip)
	if !updated {
		return errors.Wrap(errors.New("the IP is the same"), "api setter")
	}
	_, err := a.http.Get(a.generateURL(records))
	if err != nil {
		return errors.Wrap(err, "api setter")
	}
	return nil
}

// newAPISetter creates a new api setter object
func newAPISetter(c *Config, http httpClient) *apiSetter {
	return &apiSetter{config: c, http: http}
}
