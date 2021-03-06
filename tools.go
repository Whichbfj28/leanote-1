// Code generated by github.com/ssoor/gex. DO NOT EDIT.

// +build tools

package tools

// tool dependencies
import (
	_ "github.com/golang/mock/mockgen"
	_ "github.com/google/wire/cmd/wire" // Global Tools
	_ "github.com/ssoor/gex/cmd/gex"    // Global Tools
	_ "github.com/ssoor/implgen"
)

// If you want to use tools, please run the following command:
//  go generate ./tools.go
//
//go:generate go get -v github.com/ssoor/gex/cmd/gex
//go:generate go get -v github.com/google/wire/cmd/wire
//go:generate go build -v -o=./bin/implgen github.com/ssoor/implgen
//go:generate go build -v -o=./bin/mockgen github.com/golang/mock/mockgen
