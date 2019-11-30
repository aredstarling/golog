package golog

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

const timestampFormat = "2006-01-02T15:04:05.000000Z"

type stdoutLogger struct {
	name string
}

// CreateStdout a new logger
func CreateStdout() Logger {
	return CreateWriter(os.Stdout)
}

// CreateWriter a new logegr
func CreateWriter(output io.Writer) Logger {
	log.SetOutput(output)
	log.SetFlags(0)

	return &stdoutLogger{name: filepath.Base(os.Args[0])}
}

// Debug logger level logging
func (l *stdoutLogger) Debug(message string, attributes Attributes) {
	l.printLine('D', message, attributes)
}

// Error logger level logging
func (l *stdoutLogger) Error(message string, err error) {
	l.printLine('E', message, Attributes{"error": err})
}

// Fatal logger level logging
func (l *stdoutLogger) Fatal(message string, attributes Attributes) {
	l.print('F', message, attributes, log.Fatalln)
}

// FatalError logger level logging
func (l *stdoutLogger) FatalError(message string, err error) {
	l.print('F', message, Attributes{"error": err}, log.Fatalln)
}

// Info logger level logging
func (l *stdoutLogger) Info(message string, attributes Attributes) {
	l.printLine('I', message, attributes)
}

// Warn logger level logging
func (l *stdoutLogger) Warn(message string, attributes Attributes) {
	l.printLine('W', message, attributes)
}

func (l *stdoutLogger) preamble(level rune) string {
	return fmt.Sprintf("%c %s [%s:%d]:", level, time.Now().UTC().Format(timestampFormat), l.name, os.Getpid())
}

func (l *stdoutLogger) printLine(level rune, message string, attributes Attributes) {
	l.print(level, message, attributes, log.Println)
}

func (l *stdoutLogger) print(level rune, message string, attributes Attributes, print func(v ...interface{})) {
	attributes["message"] = message
	print(l.preamble(level), convert(attributes))
}

func convert(attributes Attributes) string {
	parts := make([]string, 0, len(attributes))

	for k, v := range attributes {
		sv := fmt.Sprintf("%v", v)
		parts = append(parts, fmt.Sprintf(`%s="%s"`, k, sv))
	}

	sort.Strings(parts)

	return strings.Join(parts, " ")
}
