# syntax = docker/dockerfile:1.2

# MULTI-STAGE Dockerfile

#
# go builder
#
FROM golang:1.19 as go_builder

WORKDIR /app

# Install grpcurl
RUN go install github.com/fullstorydev/grpcurl/cmd/grpcurl@v1.8.7

# Install protobuf tools, including protoc (necessary to compile protos)
RUN apt-get update \
 && DEBIAN_FRONTEND=noninteractive \
    apt-get install --no-install-recommends --assume-yes \
      protobuf-compiler

# Install protoc-gen-go (necessary to compile protos)
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
# Install protoc-gen-go (necessary to compile protos)
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

RUN mkdir proto
# Move proto files to proto folder
COPY proto ./proto

# Generate proto files
RUN mkdir bin
COPY app/bin bin
RUN sh ./bin/generate_protos.sh

# Copy all files in this directory to the /app WORKDIR, so that the container has access to all relevant code
# Leave this at the bottom always, in order to improve docker automatic caching
COPY app/ .

# Run go mod tidy to install all necessary go dependencies in the container
RUN go mod tidy

# Copy env so that the environment variables can be processed by the go code
COPY .env .

#
# go webserver code
#
FROM go_builder as go_webserver
RUN mkdir db_migrations
# Set DB migrations for this microservice
COPY postgres/hello_world/db_migrations ./db_migrations

#
# go crypto code
#
FROM go_builder as go_crypto
RUN mkdir db_migrations
# Set DB migrations for this microservice
COPY postgres/crypto/db_migrations ./db_migrations


#
# python functional tests container
#
FROM python:3.11.0 as python_tests

WORKDIR /app

COPY test/requirements.txt requirements.txt
RUN pip3 install -r requirements.txt

# Copy all files in this directory to the /app WORKDIR, so that the container has access to all relevant code
# Leave this at the bottom always, in order to improve docker automatic caching
COPY test/ .
COPY .env .


#
# nginx container
#
FROM nginx:1.23 as nginx_custom

# These lines modify the default configuration and html files for nginx
COPY ./nginx/nginx.conf /etc/nginx/conf.d/default.conf
COPY ./nginx/html /usr/share/nginx/html
