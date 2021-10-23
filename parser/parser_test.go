package parser

import (
	"github.com/stretchr/testify/assert"
	"monkey_interpreter/ast"
	"monkey_interpreter/lexer"
	"strconv"
	"testing"
)

func TestLetStatements(t *testing.T) {
	in := `
		let x = 5;
		let y = 10;
		let foobar = 83883;
	`
	l := lexer.New(in)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)

	assert.NotNil(t, program, "ParseProgram() returned nil")
	assert.Equal(t, 3, len(program.Statements))

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}
	for i, test := range tests {
		stmt := program.Statements[i]
		assert.Equal(t, "let", stmt.TokenLiteral())

		letStmt, ok := stmt.(*ast.LetStatement)
		assert.True(t, ok)
		assert.Equal(t, test.expectedIdentifier, letStmt.Name.Value)
		assert.Equal(t, test.expectedIdentifier, letStmt.Name.TokenLiteral())
	}
}

func TestReturnStatements(t *testing.T) {
	in := `
		return 5;
		return 10;
		return 83883;
	`
	l := lexer.New(in)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)

	assert.NotNil(t, program, "ParseProgram() returned nil")
	assert.Equal(t, 3, len(program.Statements))

	for _, stmt := range program.Statements {
		assert.Equal(t, "return", stmt.TokenLiteral())

		returnStmt, ok := stmt.(*ast.ReturnStatement)
		assert.True(t, ok)
		assert.Equal(t, "return", returnStmt.TokenLiteral())
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
	testIdentifier(t, consequence.Expression, "x")

	alternative, ok := ifStmt.Alternative.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)
	testIdentifier(t, alternative.Expression, "y")
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
