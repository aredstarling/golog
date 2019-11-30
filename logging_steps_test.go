// +build feature

package golog

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/DATA-DOG/godog"
)

var writer *bytes.Buffer
var logger Logger

func systemLogsADebuggingMessage(message string) error {
	logger.Debug(message, Attributes{})
	return nil
}

func systemLogsAnErrorMessageWithError(message, err string) error {
	logger.Error(message, fmt.Errorf(err))
	return nil
}

func systemLogsAnInformationMessage(message string) error {
	logger.Info(message, Attributes{})
	return nil
}

func systemLogsAWarningMessage(message string) error {
	logger.Warn(message, Attributes{})
	return nil
}

func logContainsMessage(message string) error {
	logs := strings.TrimSpace(writer.String())
	message = fmt.Sprintf(`message="%s"`, message)

	if !strings.Contains(logs, message) {
		return fmt.Errorf("Expected '%s' to contain '%s'", logs, message)
	}

	return nil
}

func logContainsMessageAndError(message, err string) error {
	if containsMessageErr := logContainsMessage(message); containsMessageErr != nil {
		return containsMessageErr
	}

	logs := strings.TrimSpace(writer.String())
	message = fmt.Sprintf(`error="%s"`, err)

	if !strings.Contains(logs, message) {
		return fmt.Errorf("Expected '%s' to contain '%s'", logs, message)
	}

	return nil
}

func setUp(interface{}) {
	writer = bytes.NewBufferString("")
	logger = CreateWriter(writer)
}

func LoggingContext(s *godog.Suite) {
	s.BeforeScenario(setUp)

	s.Step(`^the system logs a debugging message "([^"]*)"$`, systemLogsADebuggingMessage)
	s.Step(`^the system logs an error message "([^"]*)" with error "([^"]*)"$`, systemLogsAnErrorMessageWithError)
	s.Step(`^the system logs an information message "([^"]*)"$`, systemLogsAnInformationMessage)
	s.Step(`^the system logs a warning message "([^"]*)"$`, systemLogsAWarningMessage)
	s.Step(`^I should see the log contains message "([^"]*)" with error "([^"]*)"$`, logContainsMessageAndError)
	s.Step(`^I should see the log contains message "([^"]*)"$`, logContainsMessage)
}
