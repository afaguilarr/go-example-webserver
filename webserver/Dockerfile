FROM golang:latest
WORKDIR /app
COPY . .
# Install go staticcheck
RUN go install honnef.co/go/tools/cmd/staticcheck@latest
# Install dependency to save junit reports after executing unit tests
RUN go get -u github.com/jstemmer/go-junit-report
