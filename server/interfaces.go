package server

// Logger logs the information
type logger interface {
	Infof(format string, args ...interface{})
	Error(args ...interface{})
}

// Changer change the DNS records
type Changer interface {
	Change(address string)
}
