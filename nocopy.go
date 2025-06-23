package fasthttp

import "sync"

// this is a subset of the internal pragma package from github.com/protocolbuffers/protobuf-go
// see https://github.com/protocolbuffers/protobuf-go/blob/32018e9a48fe1fb3addef57cd8b745ade54fc0d7/internal/pragma/pragma.go#L29
// This does not rely on a Go language feature, but rather a special case
// within the vet checker.
//
// See https://golang.org/issues/8005.
type noCopy [0]sync.Mutex
