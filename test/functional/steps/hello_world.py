"""Module to include the step definitions related to the hello world feature"""
# Third Party
from pytest_bdd import when, then
from requests import get

# First Party
from helpers import assert_equals
from functional.steps_context import StepsContext


@when('the user requests the hello world endpoint')
def the_user_makes_an_example_api_call(steps_context: StepsContext):
    """Function to call the hello world endpoint"""
    steps_context.request = get("http://webserver:8080")


@when('the user requests the hello world endpoint with a trailing slash')
def the_user_makes_an_example_api_call_trailing_slash(steps_context: StepsContext):
    """Function to call the hello world endpoint with a trailing slash"""
    steps_context.request = get("http://webserver:8080/")


@then('the response body contains a hello world message')
def the_response_body_contains_the_expected_data(steps_context: StepsContext):
    """Function to verify the 'Hello world' message returned by the hello world endpoint"""
    body = steps_context.request.text
    assert_equals("Hello world", body)
