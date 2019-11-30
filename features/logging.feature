Feature: Logging different formats

  Make sure that we log all formats

  Scenario: Debugging message
    When the system logs a debugging message "This is a test"
    Then I should see the log contains message "This is a test"

  Scenario: Error message
    When the system logs an error message "This is a test" with error "An error!"
    Then I should see the log contains message "This is a test" with error "An error!"

  Scenario: Information message
    When the system logs an information message "This is a test"
    Then I should see the log contains message "This is a test"

  Scenario: Warning message
    When the system logs a warning message "This is a test"
    Then I should see the log contains message "This is a test"
