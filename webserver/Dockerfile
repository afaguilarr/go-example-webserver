FROM golang:latest
WORKDIR /app
COPY . .
# Install go staticcheck
RUN go install honnef.co/go/tools/cmd/staticcheck@latest
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go get -u github.com/lib/pq
