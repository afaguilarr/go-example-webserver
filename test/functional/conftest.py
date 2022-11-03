"""Module to add the common fixtures for functional tests"""
import pytest
from sqlalchemy import create_engine
from sqlalchemy.engine.base import Engine
from functional.steps_context import StepsContext
# import proto.crypto_pb2
# This will let us access the protobuf messages, such as proto.crypto_pb2.EncryptRequest

@pytest.fixture(scope="function")
def steps_context() -> StepsContext:
    """Function to make the StepsContext object available for all steps"""
    return StepsContext()


@pytest.fixture(scope="function", autouse=True)
def session() -> Engine:
    """Function to create the base state of the postgres DB"""
    db_string = "postgresql://admin:admin@postgres:5432/hello_world"
    db_session = create_engine(db_string)
    db_session.execute("DELETE FROM hello_world")
    return db_session
