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
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}
	for _, test := range tests {
		l := lexer.New(test.input)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		eval := Eval(program, env)

		testIntegerObject(t, eval, test.expected)
	}
}

func testIntegerObject(t *testing.T, obj object.Object, exp int64) {
	res, ok := obj.(*object.Integer)
	assert.True(t, ok)
	assert.Equal(t, exp, res.Value)
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != false", true},
		{"false != true", true},
		{"(1 < 2) == true", true},
		{"(1 < 2) == false", false},
		{"(1 > 2) == true", false},
		{"(1 > 2) == false", true},
	}
	for _, test := range tests {
		l := lexer.New(test.input)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		eval := Eval(program, env)

		testBooleanObject(t, eval, test.expected)
	}
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!!true", true},
		{"!!false", false},
		{"!5", false},
		{"!!5", true},
	}
	for _, test := range tests {
		l := lexer.New(test.input)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		eval := Eval(program, env)

		testBooleanObject(t, eval, test.expected)
	}
}

func testBooleanObject(t *testing.T, obj object.Object, exp bool) {
	res, ok := obj.(*object.Boolean)
	assert.True(t, ok)
	assert.Equal(t, exp, res.Value)
}

func TestIfElseExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"if (true) { 10 }", 10},
		{"if (false) { 10 }", nil},
		{"if (1) { 10 }", 10},
		{"if (1 < 2) { 10 }", 10},
		{"if (1 > 2) { 10 }", nil},
		{"if (1 > 2) { 10 } else { 20 }", 20},
		{"if (1 < 2) { 10 } else { 20 }", 10},
	}
	for _, test := range tests {
		l := lexer.New(test.input)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		eval := Eval(program, env)

		integer, ok := eval.(*object.Integer)
		if ok {
			testIntegerObject(t, integer, int64(test.expected.(int)))
		} else {
			testNullObject(t, eval)
		}
	}
}

func testNullObject(t *testing.T, obj object.Object) {
	assert.True(t, obj == NULL)
}

func TestReturnExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"return 10;", 10},
		{"return 10; 9;", 10},
		{"return 2 * 5; 9;", 10},
		{"9; return 2 * 5; 9;", 10},
		{
			`
				if(10 > 1){
					if(10 > 1) {
						return 10;					
					}
				}
				return 1;`, 10,
		},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		eval := Eval(program, env)

		testIntegerObject(t, eval, 10)
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input  string
		expMsg string
	}{
		{
			"5 + true;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"5 + true; 5;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"-true",
			"unknown operator: -BOOLEAN",
		},
		{
			"true + false;",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"5; true + false; 5",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"if (10 > 1) { true + false; }",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			`if (10 > 1) {
				if (10 > 1) {
					return true + false;
				}
				return 1;
			}`,
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"foobar;",
			"identifier not found: foobar",
		},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		eval := Eval(program, env)

		errObj, ok := eval.(*object.Error)
		assert.True(t, ok)
		assert.Equal(t, test.expMsg, errObj.Message)
	}
}

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input string
		exp   int64
	}{
		{"let a = 5; a;", 5},
		{"let a = 5 * 5; a;", 25},
		{"let a = 5; let b = a; b;", 5},
		{"let a = 5; let b = a; let c = a + b + 5; c;", 15},
	}
	for _, test := range tests {
		l := lexer.New(test.input)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		eval := Eval(program, env)

		testIntegerObject(t, eval, test.exp)
	}
}

func TestFunctionObject(t *testing.T) {
	input := "fn(x) { x + 2; };"
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	eval := Eval(program, env)

	fn, ok := eval.(*object.Function)
	assert.True(t, ok)

	assert.Equal(t, 1, len(fn.Parameters))
	assert.Equal(t, "x", fn.Parameters[0].String())
	assert.Equal(t, "(x + 2)", fn.Body.String())
}

func TestFunctionApplication(t *testing.T) {
	tests := []struct {
		input string
		exp   int64
	}{
		{"let identity = fn(x) { x; }; identity(5);", 5},
		{"let identity = fn(x) { return x; }; identity(5);", 5},
		{"let double = fn(x) { x * 2; }; double(5);", 10},
		{"let add = fn(x, y) { x + y; }; add(5, 5);", 10},
		{"let add = fn(x, y) { x + y; }; add(5 + 5, add(5, 5));", 20},
		{"fn(x) { x; }(5)", 5},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		eval := Eval(program, env)

		testIntegerObject(t, eval, test.exp)
	}
}

func TestStringLiteral(t *testing.T) {
	input := `"Hello, World!"`
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	eval := Eval(program, env)

	strObj, ok := eval.(*object.String)
	assert.True(t, ok)

	assert.Equal(t, "Hello, World!", strObj.Value)
}

func TestStringConcatenation(t *testing.T) {
	input := `"Hello" + ", " + "World!"`
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	eval := Eval(program, env)

	strObj, ok := eval.(*object.String)
	assert.True(t, ok)

	assert.Equal(t, "Hello, World!", strObj.Value)
}

func TestStringConcatenationErrorHandling(t *testing.T) {
	input := `"Hello" - ", " - "World!""`
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	eval := Eval(program, env)

	errObj, ok := eval.(*object.Error)
	assert.True(t, ok)

	assert.Equal(t, "unknown operator: STRING - STRING", errObj.Message)
}

func TestBuiltInFunctions(t *testing.T) {
	tests := []struct {
		input string
		exp   interface{}
	}{
		{`len("")`, 0},
		{`len("four")`, 4},
		{`len("hello world")`, 11},
		{`len(1)`, "argument to `len` not supported, got INTEGER"},
		{`len("one", "two")`, "wrong number of arguments. got=2, want=1"},
	}
	for _, test := range tests {
		l := lexer.New(test.input)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		eval := Eval(program, env)

		switch expected := test.exp.(type) {
		case int:
			testIntegerObject(t, eval, int64(expected))
		case string:
			errObj, ok := eval.(*object.Error)
			assert.True(t, ok)
			assert.Equal(t, test.exp, errObj.Message)
		}
	}
}

func TestArrayLiterals(t *testing.T) {
	input := `[1, 2 + 2, 3 * 3, 4 / 2]`
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	eval := Eval(program, env)

	arrObj, ok := eval.(*object.Array)
	assert.True(t, ok)

	assert.Equal(t, 4, len(arrObj.Elements))
	testIntegerObject(t, arrObj.Elements[0], 1)
	testIntegerObject(t, arrObj.Elements[1], 4)
	testIntegerObject(t, arrObj.Elements[2], 9)
	testIntegerObject(t, arrObj.Elements[3], 2)
}

func TestArrayIndexExpression(t *testing.T) {
	tests := []struct {
		input string
		exp   int
	}{
		{
			"[0, 2, 3][0]",
			0,
		},
		{
			"[0, 2, 3][1]",
			2,
		},
		{
			"[0, 2 + 2, 3 * 3][2]",
			9,
		},
		{
			"let myArray = [1, 2 * 2, 3]; myArray[0] + myArray[1] + myArray[2];",
			8,
		},
		{
			"let myArray = [1, 2, 3]; let i = myArray[0]; myArray[i]",
			2,
		},
	}
	for _, test := range tests {
		l := lexer.New(test.input)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		eval := Eval(program, env)

		testIntegerObject(t, eval, int64(test.exp))
	}
}
