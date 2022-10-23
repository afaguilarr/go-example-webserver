"""Helpers module including all base URLs"""
# Standard library
import os

# Third Party
from dotenv import load_dotenv

load_dotenv()

webserver_name = os.getenv('DOCKER_WEBSERVER_NAME')
WEBSERVER_BASE_URL = f"http://{webserver_name}:8080"
WEBSERVER_BASE_URL_TRAILING_SLASH = WEBSERVER_BASE_URL + "/"
HELLO_NAME_ENDPOINT = WEBSERVER_BASE_URL_TRAILING_SLASH + "name"
HELLO_NAME_ENDPOINT_PARAM = HELLO_NAME_ENDPOINT + "/{}"
