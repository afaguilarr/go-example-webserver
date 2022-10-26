//go:build tools
// +build tools

package tools

// This file aims to import all go tools (executable packages, not libraries being used in the code),
// After adding a dependency here, execute 'go get' for the dependency and also 'go mod tidy'
import (
	// Linting dependencies
	_ "honnef.co/go/tools/cmd/staticcheck"

	// goose for db migrations management
	_ "github.com/pressly/goose/v3/cmd/goose"

	// proto dependencies
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
