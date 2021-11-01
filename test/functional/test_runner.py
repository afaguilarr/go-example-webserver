"""Module containing the invocation of all tests derived from the feature files"""
# Third Party
from pytest_bdd import scenarios

# First Party
from functional.steps import *  # needed import for pytest-bdd pylint: disable=W0614 disable=W0401

# run all tests
scenarios('features')
