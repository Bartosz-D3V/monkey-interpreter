package evaluator

import (
	"github.com/stretchr/testify/assert"
	"monkey_interpreter/lexer"
	"monkey_interpreter/object"
	"monkey_interpreter/parser"
	"testing"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
	}
	for _, test := range tests {
		l := lexer.New(test.input)
		p := parser.New(l)
		program := p.ParseProgram()
		eval := Eval(program)

		testIntegerObject(t, eval, test.expected)
	}
}

func testIntegerObject(t *testing.T, obj object.Object, exp int64) {
	res, ok := obj.(*object.Integer)
	assert.True(t, ok)
	assert.Equal(t, exp, res.Value)
}
