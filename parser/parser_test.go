package parser

import (
	"github.com/stretchr/testify/assert"
	"monkey_interpreter/ast"
	"monkey_interpreter/lexer"
	"strconv"
	"testing"
)

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input         string
		expIdentifier string
		expValue      interface{}
	}{
		{"let x = 5;", "x", 5},
		{"let y = 10;", "y", 10},
		{"let foobar = true;", "foobar", true},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		program := p.ParseProgram()
		checkParseErrors(t, p)

		assert.NotNil(t, program, "ParseProgram() returned nil")
		assert.Equal(t, 1, len(program.Statements))
		letStatement := program.Statements[0].(*ast.LetStatement)
		assert.Equal(t, "let", letStatement.TokenLiteral())
		assert.Equal(t, test.expIdentifier, letStatement.Name.TokenLiteral())

		val := letStatement.Value
		testLiteralExpression(t, val, test.expValue)
	}
}

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input    string
		expValue interface{}
	}{
		{"return 5;", 5},
		{"return 10;", 10},
		{"return true;", true},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		program := p.ParseProgram()
		checkParseErrors(t, p)

		assert.NotNil(t, program, "ParseProgram() returned nil")
		assert.Equal(t, 1, len(program.Statements))
		returnStatement := program.Statements[0].(*ast.ReturnStatement)
		assert.Equal(t, "return", returnStatement.TokenLiteral())
		testLiteralExpression(t, returnStatement.ReturnValue, test.expValue)
	}
}

func TestIdentifier_Expression(t *testing.T) {
	in := "foobar;"
	l := lexer.New(in)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)

	assert.NotNil(t, program, "ParseProgram() returned nil")
	assert.Equal(t, 1, len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)

	identStmt, ok := stmt.Expression.(*ast.Identifier)
	assert.True(t, ok)

	assert.Equal(t, "foobar", identStmt.Value)
	assert.Equal(t, "foobar", identStmt.TokenLiteral())
}

func TestBoolean_Expression(t *testing.T) {
	in := "true;"
	l := lexer.New(in)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)

	assert.NotNil(t, program, "ParseProgram() returned nil")
	assert.Equal(t, 1, len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)

	testLiteralExpression(t, stmt.Expression, true)
}

func checkParseErrors(t *testing.T, p *Parser) {
	if len(p.errors) == 0 {
		return
	}
	t.Errorf("parser found %d error(s)", len(p.errors))
	for _, err := range p.errors {
		t.Errorf("parse error: %s", err)
	}
	t.FailNow()
}

func TestIntegerLiteral_Expression(t *testing.T) {
	in := "5;"
	l := lexer.New(in)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)

	assert.NotNil(t, program, "ParseProgram() returned nil")
	assert.Equal(t, 1, len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)

	identStmt, ok := stmt.Expression.(*ast.IntegerLiteral)
	assert.True(t, ok)

	assert.Equal(t, int64(5), identStmt.Value)
	assert.Equal(t, "5", identStmt.TokenLiteral())
}

func TestStringLiteral_Expression(t *testing.T) {
	in := `"Hello, World!"`
	l := lexer.New(in)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)

	assert.NotNil(t, program, "ParseProgram() returned nil")
	assert.Equal(t, 1, len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)

	str, ok := stmt.Expression.(*ast.StringLiteral)
	assert.True(t, ok)

	assert.Equal(t, "Hello, World!", str.Value)
	assert.Equal(t, "Hello, World!", str.TokenLiteral())
}

func TestParsePrefixExpression(t *testing.T) {
	prefixTests := []struct {
		input    string
		operator string
		val      int64
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
	}
	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParseErrors(t, p)

		assert.NotNil(t, program, "ParseProgram() returned nil")
		assert.Equal(t, 1, len(program.Statements))

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		assert.True(t, ok)

		prefixExp, ok := stmt.Expression.(*ast.PrefixExpression)
		assert.True(t, ok)

		assert.Equal(t, tt.operator, prefixExp.Operator)
		testLiteralExpression(t, prefixExp.Right, tt.val)
	}
}

func TestParseInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input    string
		leftVal  interface{}
		op       string
		rightVal interface{}
	}{
		{"5 + 5", 5, "+", 5},
		{"5 - 5", 5, "-", 5},
		{"5 * 5", 5, "*", 5},
		{"5 / 5", 5, "/", 5},
		{"5 > 5", 5, ">", 5},
		{"5 < 5", 5, "<", 5},
		{"5 == 5", 5, "==", 5},
		{"5 != 5", 5, "!=", 5},
		{"true == true", true, "==", true},
		{"false == false", false, "==", false},
		{"false != true", false, "!=", true},
	}

	for _, tt := range infixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParseErrors(t, p)

		assert.NotNil(t, program, "ParseProgram() returned nil")
		assert.Equal(t, 1, len(program.Statements))

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		assert.True(t, ok)

		infixExp, ok := stmt.Expression.(*ast.InfixExpression)
		assert.True(t, ok)

		testLiteralExpression(t, infixExp.LeftValue, tt.leftVal)
		assert.Equal(t, tt.op, infixExp.Operator)
		testLiteralExpression(t, infixExp.RightValue, tt.rightVal)
	}
}

func TestParseIfExpression(t *testing.T) {
	input := `if (x < y) { x }`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)

	assert.NotNil(t, program, "ParseProgram() returned nil")
	assert.Equal(t, 1, len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)

	ifStmt, ok := stmt.Expression.(*ast.IfExpression)
	assert.True(t, ok)

	testInfixExpression(t, ifStmt.Condition, "x", "<", "y")

	assert.Equal(t, 1, len(ifStmt.Consequence.Statements))

	consequence, ok := ifStmt.Consequence.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)
	testIdentifier(t, consequence.Expression, "x")

	assert.Nil(t, ifStmt.Alternative)
}

func TestParseIfElseExpression(t *testing.T) {
	input := `if (x < y) { x } else { y }`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)

	assert.NotNil(t, program, "ParseProgram() returned nil")
	assert.Equal(t, 1, len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)

	ifStmt, ok := stmt.Expression.(*ast.IfExpression)
	assert.True(t, ok)

	testInfixExpression(t, ifStmt.Condition, "x", "<", "y")

	assert.Equal(t, 1, len(ifStmt.Consequence.Statements))

	consequence, ok := ifStmt.Consequence.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)
	testLiteralExpression(t, consequence.Expression, "x")

	alternative, ok := ifStmt.Alternative.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)
	testLiteralExpression(t, alternative.Expression, "y")
}

func TestParseFunctionLiteral(t *testing.T) {
	input := `fn(x, y) { x + y; }`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)

	assert.NotNil(t, program, "ParseProgram() returned nil")
	assert.Equal(t, 1, len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)

	fnStmt, ok := stmt.Expression.(*ast.FunctionLiteral)
	assert.True(t, ok)

	assert.Equal(t, 2, len(fnStmt.Parameters))
	testLiteralExpression(t, fnStmt.Parameters[0], "x")
	testLiteralExpression(t, fnStmt.Parameters[1], "y")

	assert.Equal(t, 1, len(fnStmt.Body.Statements))
	bodyExp, ok := fnStmt.Body.Statements[0].(*ast.ExpressionStatement)
	testInfixExpression(t, bodyExp.Expression, "x", "+", "y")
}

func TestParseFunctionLiteralParams(t *testing.T) {
	fnTests := []struct {
		input     string
		expParams []string
	}{
		{"fn(x,y){}", []string{"x", "y"}},
		{"fn(x){}", []string{"x"}},
		{"fn(){}", []string{}},
	}
	for _, tt := range fnTests {
		input := tt.input

		l := lexer.New(input)
		p := New(l)
		program := p.ParseProgram()
		checkParseErrors(t, p)

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		assert.True(t, ok)

		fnStmt, ok := stmt.Expression.(*ast.FunctionLiteral)
		assert.True(t, ok)

		assert.Equal(t, len(tt.expParams), len(fnStmt.Parameters))
		for i, expParam := range tt.expParams {
			testLiteralExpression(t, fnStmt.Parameters[i], expParam)
		}
	}
}

func TestCallExpressionParsing(t *testing.T) {
	input := `add(1, 2 + 3, 4 * 8);`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)

	assert.NotNil(t, program, "ParseProgram() returned nil")
	assert.Equal(t, 1, len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)

	callExp, ok := stmt.Expression.(*ast.CallExpression)
	assert.True(t, ok)

	testIdentifier(t, callExp.Function, "add")
	assert.Equal(t, 3, len(callExp.Arguments))
	testLiteralExpression(t, callExp.Arguments[0], 1)
	testInfixExpression(t, callExp.Arguments[1], 2, "+", 3)
	testInfixExpression(t, callExp.Arguments[2], 4, "*", 8)
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"-a * b",
			"((-a) * b)",
		},
		{
			"-a * -b",
			"((-a) * (-b))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
		{
			"true",
			"true",
		},
		{
			"false",
			"false",
		},
		{
			"3 > 2 == true",
			"((3 > 2) == true)",
		},
		{
			"3 < 5 == false",
			"((3 < 5) == false)",
		},
		{
			"1 + (2 + 3) == 6",
			"((1 + (2 + 3)) == 6)",
		},
		{
			"(5 + 5) * 2 == 20",
			"(((5 + 5) * 2) == 20)",
		},
		{
			"-(5 + 5) * 2 == -20",
			"(((-(5 + 5)) * 2) == (-20))",
		},
		{
			"!(true == true)",
			"(!(true == true))",
		},
		{
			"a + add(b * c) + d",
			"((a + add((b * c))) + d)",
		},
		{
			"add(a, b, 1, 2 * 3, 4 + 5, add(6, 7 * 8))",
			"add(a, b, 1, (2 * 3), (4 + 5), add(6, (7 * 8)))",
		},
		{
			"add(a + b + c * d / f + g)",
			"add((((a + b) + ((c * d) / f)) + g))",
		},
		{
			"a * [1, 2, 3, 4][b * c] * d",
			"((a * ([1, 2, 3, 4][(b * c)])) * d)",
		},
		{
			"add(a * b[2], b[1], 2 * [1, 2][1])",
			"add((a * (b[2])), (b[1]), (2 * ([1, 2][1])))",
		},
	}
	for i, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParseErrors(t, p)
		actual := program.String()

		assert.Equalf(t, tt.expected, actual, "parsing statement with index %d failed", i)
	}
}

func TestArrayLiteral(t *testing.T) {
	input := `[1, 2 * 2, 3 - 2]`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)

	assert.NotNil(t, program, "ParseProgram() returned nil")
	assert.Equal(t, 1, len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)

	arrLit, ok := stmt.Expression.(*ast.ArrayLiteral)
	assert.True(t, ok)
	testLiteralExpression(t, arrLit.Elements[0], 1)
	testInfixExpression(t, arrLit.Elements[1], 2, "*", 2)
	testInfixExpression(t, arrLit.Elements[2], 3, "-", 2)
}

func TestIndexExpression(t *testing.T) {
	input := `myArray[1 + 2]`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)

	assert.NotNil(t, program, "ParseProgram() returned nil")
	assert.Equal(t, 1, len(program.Statements))

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)

	idxExp, ok := stmt.Expression.(*ast.IndexExpression)
	testIdentifier(t, idxExp.Left, "myArray")
	testInfixExpression(t, idxExp.Index, 1, "+", 2)
}

func testLiteralExpression(t *testing.T, exp ast.Expression, expected interface{}) {
	switch v := expected.(type) {
	case int:
	case int64:
		testIntegerLiteral(t, v, exp)
	case string:
		testIdentifier(t, exp, v)
	case bool:
		testBoolean(t, exp, v)
	}
}

func testInfixExpression(t *testing.T, exp ast.Expression, left interface{}, op string, right interface{}) {
	opExp, ok := exp.(*ast.InfixExpression)
	assert.True(t, ok)

	testLiteralExpression(t, opExp.LeftValue, left)
	testLiteralExpression(t, opExp.RightValue, right)
	assert.Equal(t, opExp.Operator, op)
}

func testIntegerLiteral(t *testing.T, val int64, right ast.Expression) {
	intLit, ok := right.(*ast.IntegerLiteral)
	assert.True(t, ok)

	assert.Equal(t, val, intLit.Value)
	assert.Equal(t, strconv.FormatInt(val, 10), intLit.TokenLiteral())
}

func testIdentifier(t *testing.T, exp ast.Expression, expected string) {
	ident, ok := exp.(*ast.Identifier)
	assert.True(t, ok)

	assert.Equal(t, expected, ident.Value)
	assert.Equal(t, expected, ident.TokenLiteral())
}

func testBoolean(t *testing.T, exp ast.Expression, expected bool) {
	ident, ok := exp.(*ast.Boolean)
	assert.True(t, ok)

	assert.Equal(t, expected, ident.Value)
	assert.Equal(t, strconv.FormatBool(expected), ident.TokenLiteral())
}
