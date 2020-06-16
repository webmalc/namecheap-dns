package cmd

// ErrorLogger logs errors.
type ErrorLogger interface {
	Error(args ...interface{})
}

// Changer change the DNS
type Changer interface {
	Change(address string)
}

// Server runs the server
type Server interface {
	Run()
}

// Client sends the request
type Client interface {
	Request()
}
