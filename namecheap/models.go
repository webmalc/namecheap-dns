package namecheap

// apiResponse is the api response struct.
type apiResponse struct {
	Status string `xml:"Status,attr"`
	Hosts  []host `xml:"CommandResponse>DomainDNSGetHostsResult>host"`
}

// host is the host record struct.
type host struct {
	ID                 int    `xml:"HostId,attr"`
	Name               string `xml:"Name,attr"`
	Type               string `xml:"Type,attr"`
	Address            string `xml:"Address,attr"`
	MXPref             int    `xml:"MXPref,attr"`
	TTL                int    `xml:"TTL,attr"`
	AssociatedAppTitle string `xml:"AssociatedAppTitle,attr"`
	FriendlyName       string `xml:"FriendlyName,attr"`
	IsActive           bool   `xml:"IsActive,attr"`
	IsDDNSEnabled      bool   `xml:"IsDDNSEnabled,attr"`
}
