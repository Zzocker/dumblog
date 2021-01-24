package dumblog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockWriter struct {
	stream string
}

func (m *mockWriter) Write(p []byte) (int, error) {
	m.stream = string(p)
	return len(p), nil
}

func (m *mockWriter) clean() {
	m.stream = ""
}

func TestErrNonErrStream(t *testing.T) {
	assert := assert.New(t)
	outStream := &mockWriter{}
	errStream := &mockWriter{}
	lgr := &logger{
		out:    outStream,
		errOut: errStream,
		l:      DebugLevel,
	}
	msg := "ErrNonErrStream"
	outMsg := msg + string('\n')

	// debug level msg
	lgr.output(DebugLevel, msg)
	assert.Equal(outMsg, outStream.stream)
	assert.Zero(errStream.stream)
	outStream.clean()

	// info level msg
	lgr.output(InfoLevel, msg)
	assert.Equal(outMsg, outStream.stream)
	assert.Zero(errStream.stream)
	outStream.clean()

	// warn level msg
	lgr.output(WarnLevel, msg)
	assert.Equal(outMsg, outStream.stream)
	assert.Zero(errStream.stream)
	outStream.clean()

	// error level msg
	lgr.output(ErrorLevel, msg)
	assert.Equal(outMsg, errStream.stream)
	assert.Zero(outStream.stream)
	errStream.clean()
}

// level less then Info should not stream to output
func TestInfoLevel(t *testing.T) {
	assert := assert.New(t)
	outStream := &mockWriter{}
	errStream := &mockWriter{}
	lgr := &logger{
		out:    outStream,
		errOut: errStream,
		l:      InfoLevel,
	}
	msg := "InfoLevel"
	outMsg := msg + string('\n')

	// debug level msg should not result in stream
	lgr.output(DebugLevel, msg)
	assert.Zero(outStream.stream)

	lgr.output(InfoLevel, msg)
	assert.Equal(outMsg, outStream.stream)
	outStream.clean()

	lgr.output(WarnLevel, msg)
	assert.Equal(outMsg, outStream.stream)
	outStream.clean()

	lgr.output(ErrorLevel, msg)
	assert.Equal(outMsg, errStream.stream)
	errStream.clean()
}

func TestWarnLevel(t *testing.T) {
	assert := assert.New(t)
	outStream := &mockWriter{}
	errStream := &mockWriter{}
	lgr := &logger{
		out:    outStream,
		errOut: errStream,
		l:      WarnLevel,
	}
	msg := "WarnLevel"
	outMsg := msg + string('\n')

	// debug level msg should not result in stream
	lgr.output(DebugLevel, msg)
	assert.Zero(outStream.stream)

	lgr.output(InfoLevel, msg)
	assert.Zero(outStream.stream)

	lgr.output(WarnLevel, msg)
	assert.Equal(outMsg, outStream.stream)
	outStream.clean()

	lgr.output(ErrorLevel, msg)
	assert.Equal(outMsg, errStream.stream)
	errStream.clean()
}

func TestErrorLevel(t *testing.T) {
	assert := assert.New(t)
	outStream := &mockWriter{}
	errStream := &mockWriter{}
	lgr := &logger{
		out:    outStream,
		errOut: errStream,
		l:      ErrorLevel,
	}
	msg := "ErrorLevel"
	outMsg := msg + string('\n')

	// debug level msg should not result in stream
	lgr.output(DebugLevel, msg)
	assert.Zero(outStream.stream)

	lgr.output(InfoLevel, msg)
	assert.Zero(outStream.stream)

	lgr.output(WarnLevel, msg)
	assert.Zero(outStream.stream)

	lgr.output(ErrorLevel, msg)
	assert.Equal(outMsg, errStream.stream)
	errStream.clean()
}

func TestNewLineAppend(t *testing.T) {
	assert := assert.New(t)
	outStream := &mockWriter{}
	errStream := &mockWriter{}
	lgr := &logger{
		out:    outStream,
		errOut: errStream,
		l:      DebugLevel,
	}

	msg := "WithoutNewLine"

	lgr.output(DebugLevel, msg)
	assert.Equal('\n', rune(outStream.stream[len(outStream.stream)-1]))
	outStream.clean()

	lgr.output(ErrorLevel, msg)
	assert.Equal('\n', rune(errStream.stream[len(errStream.stream)-1]))
	errStream.clean()
}
