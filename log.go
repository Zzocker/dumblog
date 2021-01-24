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

// output will write s to output streams based on lvl
// if s doesn't have \n then it will be appended at then end of s
// error lvl s will streamed to errStream
func (l *logger) output(lvl level, s string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	// if lvl is less then set level
	// then no need to stream
	if lvl < l.l {
		return
	}

	// append newline
	if len(s) == 0 || s[len(s)-1] != '\n' {
		s += string('\n')
	}
	if lvl == ErrorLevel {
		l.errOut.Write([]byte(s))
	} else {
		l.out.Write([]byte(s))
	}
}
