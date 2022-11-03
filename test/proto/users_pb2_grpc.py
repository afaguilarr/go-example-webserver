# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from proto import users_pb2 as proto_dot_users__pb2


class UsersStub(object):
    """Users exposes all the necessary RPCs to handle Users information.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Register = channel.unary_unary(
                '/go_webserver.users.Users/Register',
                request_serializer=proto_dot_users__pb2.RegisterRequest.SerializeToString,
                response_deserializer=proto_dot_users__pb2.RegisterResponse.FromString,
                )


class UsersServicer(object):
    """Users exposes all the necessary RPCs to handle Users information.
    """

    def Register(self, request, context):
        """Register receives all necessary information to create a new user,
        and then returns the stored information.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_UsersServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Register': grpc.unary_unary_rpc_method_handler(
                    servicer.Register,
                    request_deserializer=proto_dot_users__pb2.RegisterRequest.FromString,
                    response_serializer=proto_dot_users__pb2.RegisterResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'go_webserver.users.Users', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Users(object):
    """Users exposes all the necessary RPCs to handle Users information.
    """

    @staticmethod
    def Register(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/go_webserver.users.Users/Register',
            proto_dot_users__pb2.RegisterRequest.SerializeToString,
            proto_dot_users__pb2.RegisterResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
