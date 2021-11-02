"""Module to add the common fixtures for functional tests"""
import pytest
from functional.steps_context import StepsContext


@pytest.fixture(scope="function")
def steps_context() -> StepsContext:
    """Function to make the StepsContext object available for all steps"""
    return StepsContext()
