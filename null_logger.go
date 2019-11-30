package golog

type nullLogger struct {
}

// CreateNull a new logger
func CreateNull() Logger {
	return &nullLogger{}
}

// Debug logger level logging
func (l *nullLogger) Debug(message string, attributes Attributes) {
}

// Error logger level logging
func (l *nullLogger) Error(message string, err error) {
}

// Fatal logger level logging
func (l *nullLogger) Fatal(message string, attributes Attributes) {
}

// FatalError logger level logging
func (l *nullLogger) FatalError(message string, err error) {
}

// Info logger level logging
func (l *nullLogger) Info(message string, attributes Attributes) {
}

// Warn logger level logging
func (l *nullLogger) Warn(message string, attributes Attributes) {
}
