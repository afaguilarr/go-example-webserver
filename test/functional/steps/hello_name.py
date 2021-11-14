"""xD"""
# Third Party
from pytest_bdd import when, then, given
from requests import get
from sqlalchemy.engine.base import Engine

# First Party
from helpers import assert_equals, assert_not_none
from functional.steps_context import StepsContext


@given('a random name', target_fixture="name")
def random_name() -> str:
    """xD"""
    return "alfonso"


@when('the user requests the hello name endpoint')
def the_user_requests_hello_name_endpoint(name: str, steps_context: StepsContext):
    """xD"""
    steps_context.request = get(f"http://webserver:8080/name/{name}")


@then('the response body contains the hello name message')
def the_response_body_contains_the_hello_name_data(name: str, steps_context: StepsContext):
    """xD"""
    body = steps_context.request.text
    assert_equals(f"Hello {name}", body)


@then('the name was properly stored')
def name_was_properly_stored(name: str, session: Engine):
    """xD"""
    result_set = session.execute("SELECT * FROM hello_world")
    results = list(result_set)
    assert_equals(1, len(results))
    result = results[0]
    assert_not_none(result[0])
    assert_equals(name, result[1])
