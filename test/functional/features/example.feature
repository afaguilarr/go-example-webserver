@API
Feature: Esample

  Scenario: Hello World
    When the user requests the hello world endpoint
    Then the API responds successfully
    And the response body contains a hello world message
