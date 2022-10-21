FROM golang:1.19

WORKDIR /app

# Install go staticcheck
RUN go install honnef.co/go/tools/cmd/staticcheck@latest
# Install goose for DB migrations
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
# Install Go postgres driver
RUN go get -u github.com/lib/pq

COPY . .
