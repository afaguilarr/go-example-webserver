FROM golang:latest
WORKDIR /app
COPY . .
RUN rm ./go.mod
RUN go mod init goWebServer
