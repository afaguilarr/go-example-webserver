python -m grpc_tools.protoc -I/ --python_out=./ --grpc_python_out=./ /proto/common.proto
python -m grpc_tools.protoc -I/ --python_out=./ --grpc_python_out=./ /proto/crypto.proto
python -m grpc_tools.protoc -I/ --python_out=./ --grpc_python_out=./ /proto/users.proto
