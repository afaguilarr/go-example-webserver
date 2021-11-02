if `go vet ./main | grep go`; then exit 1; fi
if `go fmt ./main | grep go`; then exit 1; fi
