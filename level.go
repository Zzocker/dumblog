package dumblog

type level int8

const (
	// DebugLevel : by default debug level is off in production
	DebugLevel level = iota - 1

	// InfoLevel is deafult logging level
	InfoLevel

	// WarnLevel has more priority then info,
	// but not that much, in short it doesn't require human review
	WarnLevel

	// ErrorLevel is high-priority and should not occurred
	// defiantly require human review
	ErrorLevel

	// FatalLevel highest in priority
	// will log the message and then os.Exit(1) is called	
	FatalLevel
)
