package dumblog

import (
	"io"
	"sync"
)

type logger struct {
	out    io.Writer
	errOut io.Writer
	mu     sync.Mutex
	l      level
}

func (l *logger) Debug(...interface{}) {
	// TODO
}
func (l *logger) Debugf(string, ...interface{}) {
	// TODO
}

func (l *logger) Info(...interface{}) {
	// TODO
}
func (l *logger) Infof(string, ...interface{}) {
	// TODO
}

func (l *logger) Warn(...interface{}) {
	// TODO
}
func (l *logger) Warnf(string, ...interface{}) {
	// TODO
}

func (l *logger) Fatal(...interface{}) {
	// TODO
}
func (l *logger) Fatalf(string, ...interface{}) {
	// TODO
}
