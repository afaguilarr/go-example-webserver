from typing import Any


def assert_equals(expected: Any, actual: Any):
    assert expected == actual, f"Expected\n{expected}\nto equal\n{actual}"
