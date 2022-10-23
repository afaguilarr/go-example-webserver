"""Helpers module to gather common code related to http calls"""
# Third Party
from requests import get, Response


def get_request(url: str, timeout_in_seconds: int = 3) -> Response:
    """Helper function to send a requests.get call with a hardcoded timeout"""
    return get(url, timeout=timeout_in_seconds)
