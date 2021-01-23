package dumblog

import "io"

// DumbLog represents an active logging object
// that generates lines of output to an io.Writer
type DumbLog interface {
	Debug(...interface{})
	Debugf(string, ...interface{})

	Info(...interface{})
	Infof(string, ...interface{})

	Warn(...interface{})
	Warnf(string, ...interface{})

	Fatal(...interface{})
	Fatalf(string, ...interface{})
}

// New : creates a thread safe dumblogger
func New(level level, errOut, out io.Writer) DumbLog {
	return &logger{
		out:    out,
		errOut: errOut,
		l:      level,
	}
}
