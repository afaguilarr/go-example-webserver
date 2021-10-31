from functional.steps_context import StepsContext
from pytest_bdd import given, when, then
from helpers import assert_equals
from requests import get


@when('the user requests the hello world endpoint')
def the_user_makes_an_example_api_call(steps_context: StepsContext):
    steps_context.request = get("http://webserver:8080")


@then("the API responds successfully")
def the_api_responds_successfully(steps_context: StepsContext):
    assert steps_context.request.status_code == 200


@then('the response body contains a hello world message')
def the_response_body_contains_the_expected_data(steps_context: StepsContext):
    body = steps_context.request.text
    assert_equals("Hello world", body)
