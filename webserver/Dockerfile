FROM golang:latest
WORKDIR /app
COPY . .
# Install go staticcheck
RUN go install honnef.co/go/tools/cmd/staticcheck@latest
