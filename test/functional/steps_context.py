"""
Module containing the StepsContext class, all variables shared
between when and then steps can be added here
"""
from dataclasses import dataclass
from typing import Any


@dataclass
class StepsContext():
    """Class to add all variables shared between when and then steps"""
    request: Any = None
