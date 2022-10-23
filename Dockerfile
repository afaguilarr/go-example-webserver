# syntax = docker/dockerfile:1.2

# MULTI-STAGE Dockerfile

#
# go webserver code
#
FROM golang:1.19 as go_webserver

WORKDIR /app

# Install go staticcheck
RUN go install honnef.co/go/tools/cmd/staticcheck@latest
# Install goose for DB migrations
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Copy all files in this directory to the /app WORKDIR, so that the container has access to all relevant code
# Leave this at the bottom always, in order to improve docker automatic caching
COPY webserver/ .
COPY .env .


#
# python functional tests container
#
FROM python:3.10 as python_tests

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
