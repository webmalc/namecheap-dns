package cmd

// ErrorLogger logs errors.
type ErrorLogger interface {
	Error(args ...interface{})
}

// Changer change the DNS
type Changer interface {
	Change(address string)
}
