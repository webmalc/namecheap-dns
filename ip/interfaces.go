package ip

// Logger logs the information
type logger interface {
	Infof(format string, args ...interface{})
	Error(args ...interface{})
}
