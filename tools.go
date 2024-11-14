//go:build tools
// +build tools

package main

import (
	// https://golangci-lint.run
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	// https://go.dev/blog/vuln
	_ "golang.org/x/vuln/cmd/govulncheck"
	// gRPC generation via https://buf.build/
	_ "github.com/bufbuild/buf/cmd/buf"
	_ "go.etcd.io/bbolt"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
