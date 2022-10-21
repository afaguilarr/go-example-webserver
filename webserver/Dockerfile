FROM golang:1.19

WORKDIR /app

# Install go staticcheck
RUN go install honnef.co/go/tools/cmd/staticcheck@latest
# Install goose for DB migrations
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Copy all files in this directory to the /app WORKDIR, so that the container has access to all relevant code
# Leave this at the bottom always, in order to improve docker automatic caching
COPY . .
