import pytest
from functional.steps_context import StepsContext


@pytest.fixture(scope="function")
def steps_context() -> StepsContext:
    return StepsContext()
