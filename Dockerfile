# syntax = docker/dockerfile:1.2

# MULTI-STAGE Dockerfile

#
# go builder
#
FROM golang:1.19 as go_builder

WORKDIR /app

# Install necessary Linux tools
RUN apt-get update && \
    apt-get install --yes --quiet \
      unzip \
      jq \
    && rm --force --recursive /var/lib/apt/lists/*

# Install grpcurl
RUN go install github.com/fullstorydev/grpcurl/cmd/grpcurl@v1.8.7

# Download protoc and install to /usr/local/bin
ENV PROTOC_VERSION 3.14.0
RUN mkdir --parents "/tmp/protoc-${PROTOC_VERSION}-$(uname -s)-$(uname -m)/"
RUN curl --fail --show-error \
      --location "https://github.com/google/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-$(uname -s)-$(uname -m).zip" \
      --output "/tmp/protoc-${PROTOC_VERSION}-$(uname -s)-$(uname -m)/protoc-${PROTOC_VERSION}-$(uname -s)-$(uname -m).zip"
RUN cd "/tmp/protoc-${PROTOC_VERSION}-$(uname -s)-$(uname -m)/" && \
    unzip "protoc-${PROTOC_VERSION}-$(uname -s)-$(uname -m).zip" && \
    cp bin/protoc /usr/local/bin/protoc && \
    mkdir --parents /usr/local/include/google/protobuf && \
    # This step is really important, since it copies all the well known types where protoc will find them
    cp --recursive include/google/protobuf/* /usr/local/include/google/protobuf/ && \
    rm -rf "/tmp/protoc-${PROTOC_VERSION}-$(uname -s)-$(uname -m)/"

# Download go-swagger and install to /usr/local/bin
RUN curl -o /usr/local/bin/swagger -L'#' "$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest | jq -r '.assets[] | select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url')"
RUN chmod +x /usr/local/bin/swagger

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

# Set db_migrations folder for the child configs
RUN mkdir db_migrations

#
# go webserver code
#
FROM go_builder as go_webserver
# Set DB migrations for this microservice
COPY postgres/hello_world/db_migrations ./db_migrations

#
# go crypto code
#
FROM go_builder as go_crypto
# Set DB migrations for this microservice
COPY postgres/crypto/db_migrations ./db_migrations

#
# go users code
#
FROM go_builder as go_users
# Set DB migrations for this microservice
COPY postgres/users/db_migrations ./db_migrations


#
# python functional tests container
#
FROM python:3.11.1 as python_tests

WORKDIR /app

COPY test/requirements.txt requirements.txt
RUN pip3 install -r requirements.txt

RUN mkdir proto
RUN mkdir /proto
# Move proto files to proto folder
COPY proto/ /proto
RUN mkdir bin
COPY test/bin/ ./bin
RUN sh bin/generate_protos.sh

# Copy all files in this directory to the /app WORKDIR, so that the container has access to all relevant code
# Leave this at the bottom always, in order to improve docker automatic caching
COPY test/ .
COPY .env .


#
# nginx container
#
FROM nginx:1.23.3 as nginx_custom

# These lines modify the default configuration and html files for nginx
COPY ./nginx/nginx.conf /etc/nginx/conf.d/default.conf
COPY ./nginx/html /usr/share/nginx/html
