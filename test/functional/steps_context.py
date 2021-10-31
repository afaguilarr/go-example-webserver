from dataclasses import dataclass
from typing import Any


@dataclass
class StepsContext():
    request: Any = None
