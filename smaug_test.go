package smaug

import (
	"testing"

	"github.com/expr-lang/expr"
)

const EXPR_SCRIPT = `
IntValue
`

type TestData struct {
	IntValue int
	StrValue string
}

func TestRun(t *testing.T) {
	program, err := Compile(EXPR_SCRIPT, expr.AsInt(), expr.Env(TestData{}))
	if err != nil {
		t.Fatal("compile expected no err, but got", err)
	}

	data := TestData{
		IntValue: 2024,
		StrValue: "hello",
	}
	result, err := Run(program, data)
	if err != nil {
		t.Fatal("run expected no err, but got", err)
	}
	r, ok := result.(int)
	if !ok {
		t.Fatal("run result expected int, but not")
	}
	if r != 2024 {
		t.Fatal("run result not correct")
	}
}

