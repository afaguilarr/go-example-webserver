"""Module to add all the assertions functions"""
from typing import Any


def assert_equals(expected: Any, actual: Any):
    """Function to assert that 2 values are equal"""
    assert expected == actual, f"Expected\n{expected}\nto equal\n{actual}"


def assert_not_none(expected: Any):
    """xD"""
    assert expected is not None, f"Expected\n{expected}\nto not equal\nNone"
