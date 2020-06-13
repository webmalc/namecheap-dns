package namecheap

import (
	"net"
	"testing"

	"github.com/imroc/req"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Should update the A host records
func Test_apiSetter_updateRecords(t *testing.T) {
	c := newConfig()
	h := &httpClientMock{}
	s := newAPISetter(c, h)
	ip := "0.0.0.0"
	records := &apiResponse{
		Status: "OK",
		Hosts: []host{
			{Type: "A", Address: ip},
			{Type: "MX", Address: ip},
			{Type: "A", Address: ip},
		},
	}
	updated := s.updateRecords(records, net.ParseIP(ip))
	assert.False(t, updated)
	updated = s.updateRecords(records, net.ParseIP("1.1.1.1"))
	assert.True(t, updated)
	for _, host := range records.Hosts {
		if host.Type == "A" {
			assert.Equal(t, "1.1.1.1", host.Address)
		} else {
			assert.Equal(t, ip, host.Address)
		}
	}
}

// Should generate an URL to update the host records
func Test_apiSetter_generateURL(t *testing.T) {
	c := newConfig()
	h := &httpClientMock{}
	s := newAPISetter(c, h)
	records := &apiResponse{
		Status: "OK",
		Hosts: []host{
			{Type: "A", Address: "1.1.1.1", Name: "one", TTL: 10},
			{Type: "MX", Address: "2.2.2.2", Name: "two", TTL: 20, MXPref: 10},
			{Type: "A", Address: "3.3.3.3", Name: "three", TTL: 30},
		},
	}
	expected := "https://api.namecheap.com/xml.response?ApiUser=user" +
		"&ApiKey=key&UserName=username&ClientIp=ip&SLD=domain&TLD=com&Command" +
		"=namecheap.domains.dns.setHosts&HostName1=one&RecordType1=A" +
		"&Address1=1.1.1.1&TTL1=10&HostName2=two&RecordType2=MX&" +
		"Address2=2.2.2.2&TTL2=20&MXPref2=10&HostName3=three&" +
		"RecordType3=A&Address3=3.3.3.3&TTL3=30&"
	u := s.generateURL(records)
	assert.Equal(t, expected, u)
}

// Should return an error
func Test_apiSetter_Set_Same_IP(t *testing.T) {
	c := newConfig()
	h := &httpClientMock{}
	s := newAPISetter(c, h)
	ip := "127.0.0.1"
	records := &apiResponse{
		Status: "OK",
		Hosts:  []host{{Type: "A", Address: ip}},
	}
	err := s.Set(records, net.ParseIP(ip))
	assert.Error(t, err)
}

// Should return an error
func Test_apiSetter_Set_Invalid_Request(t *testing.T) {
	c := newConfig()
	h := &httpClientErrorMock{}
	s := newAPISetter(c, h)
	records := &apiResponse{
		Status: "OK",
		Hosts:  []host{{Type: "A", Address: "1.1.1.1"}},
	}
	h.On("Get", mock.Anything, mock.Anything).Return(
		&req.Resp{}, errors.New("test error"),
	).Once()
	err := s.Set(records, net.ParseIP("2.2.2.2"))
	assert.Error(t, err)
	h.AssertExpectations(t)
}

// Should create a new object
func Test_newAPISetter(t *testing.T) {
	c := newConfig()
	h := &httpClientMock{}
	s := newAPISetter(c, h)
	assert.Equal(t, s.config, c)
	assert.Equal(t, s.http, h)
}
