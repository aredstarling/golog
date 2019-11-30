package golog

// Logger allows to log diagnostics
type Logger interface {
	Debug(message string, attributes Attributes)
	Error(message string, err error)
	Fatal(message string, attributes Attributes)
	FatalError(message string, err error)
	Info(message string, attributes Attributes)
	Warn(message string, attributes Attributes)
}
