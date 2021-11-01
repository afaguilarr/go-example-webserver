@API
Feature: Hello World
  As a webserver user,
  I want a GET endpoint to return a simple Hello World message,
  so that I know the webserver is healthy.

  Scenario: Hello World is returned when using the root path
    When the user requests the hello world endpoint
    Then the response has a "200" status code
    And the response body contains a hello world message

  Scenario: Hello World is returned when using the root path with a trailing slash
    When the user requests the hello world endpoint with a trailing slash
    Then the response has a "200" status code
    And the response body contains a hello world message
