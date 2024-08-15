//go:build (!amd64 && !arm64) || go1.24 || !go1.17 || (arm64 && !go1.20)
// +build !amd64,!arm64 go1.24 !go1.17 arm64,!go1.20

package smaug

import (
	"github.com/expr-lang/expr"
)

var (
	Compile = expr.Compile
	Run     = expr.Run
)

