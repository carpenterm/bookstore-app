//go:build tools

// Package main's seed scripts build with "ignore", which go mod tidy skips;
// this keeps their dependencies in go.mod.
package main

import (
	_ "github.com/minio/minio-go/v7"
	_ "github.com/minio/minio-go/v7/pkg/credentials"
	_ "golang.org/x/image/font"
	_ "golang.org/x/image/font/gofont/gobold"
	_ "golang.org/x/image/font/opentype"
	_ "golang.org/x/image/math/fixed"
)
