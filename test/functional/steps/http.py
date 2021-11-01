"""Module to add the http common steps"""
# Third Party
from pytest_bdd import then, parsers

# First Party
from helpers import assert_equals
from functional.steps_context import StepsContext


@then(parsers.parse('the response has a "{status_code}" status code'))
def response_has_status_code(steps_context: StepsContext, status_code: str):
    #"""Function to check that the response contains certain status code"""
    assert_equals(int(status_code), steps_context.request.status_code)
