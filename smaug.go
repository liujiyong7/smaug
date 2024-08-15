//go:build (amd64 && go1.17 && !go1.24) || (arm64 && go1.20 && !go1.24)
// +build amd64,go1.17,!go1.24 arm64,go1.20,!go1.24

package smaug

import (
	"github.com/liujiyong7/smaug/internal/core"
	"github.com/expr-lang/expr"
	"github.com/expr-lang/expr/vm"
)

func Compile(input string, ops ...expr.Option) (*vm.Program, error) {
	program, err := expr.Compile(input, ops...)
	if err != nil {
		return nil, err
	}
	core.CompileBinaryAndStore(program)
	return program, nil
}

func Run(program *vm.Program, env any) (any, error) {
	result, err := core.FindBinaryAndRun(program, env)
	if err == nil {
		return result, nil
	}
	return vm.Run(program, env)
}

