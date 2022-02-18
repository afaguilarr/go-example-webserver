@API
Feature: Hello Name
#   As a webserver user,
#   I want a GET endpoint to return a simple Hello World message,
#   so that I know the webserver is healthy.

  Scenario: Hello name saves the expeted entities and returns the expected text
    Given a random name
    When the user requests the hello name endpoint
    Then the response has a "200" status code
    And the response body contains the hello name message
    And the name was properly stored
