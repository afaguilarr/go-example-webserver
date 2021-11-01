from functional.steps_context import StepsContext
from pytest_bdd import then, parsers
from helpers import assert_equals


@then(parsers.parse('the response has a "{status_code}" status code'))
def the_api_responds_successfully(steps_context: StepsContext, status_code: str):
    assert_equals(int(status_code), steps_context.request.status_code)
