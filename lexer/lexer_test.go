package lexer

import (
	"monkey_interpreter/token"
	"testing"
)

func TestLexer_NextToken_Signs(t *testing.T) {
	input := "=+(){},;"
	tests := []struct {
		ExpectedType    token.Type
		ExpectedLiteral string
	}{
		{ExpectedType: token.ASSIGN, ExpectedLiteral: "="},
		{ExpectedType: token.PLUS, ExpectedLiteral: "+"},
		{ExpectedType: token.LPAREN, ExpectedLiteral: "("},
		{ExpectedType: token.RPAREN, ExpectedLiteral: ")"},
		{ExpectedType: token.LBRACE, ExpectedLiteral: "{"},
		{ExpectedType: token.RBRACE, ExpectedLiteral: "}"},
		{ExpectedType: token.COMMA, ExpectedLiteral: ","},
		{ExpectedType: token.SEMICOLON, ExpectedLiteral: ";"},
		{ExpectedType: token.EOF, ExpectedLiteral: ""},
	}
	runT(t, input, tests)
}

func TestLexer_NextToken_LogicOps(t *testing.T) {
	input := `
		10 == 10;
		10 != 9;
	`
	tests := []struct {
		ExpectedType    token.Type
		ExpectedLiteral string
	}{
		{ExpectedType: token.INT, ExpectedLiteral: "10"},
		{ExpectedType: token.EQ, ExpectedLiteral: "=="},
		{ExpectedType: token.INT, ExpectedLiteral: "10"},
		{ExpectedType: token.SEMICOLON, ExpectedLiteral: ";"},

		{ExpectedType: token.INT, ExpectedLiteral: "10"},
		{ExpectedType: token.NEQ, ExpectedLiteral: "!="},
		{ExpectedType: token.INT, ExpectedLiteral: "9"},
		{ExpectedType: token.SEMICOLON, ExpectedLiteral: ";"},

		{ExpectedType: token.EOF, ExpectedLiteral: ""},
	}
	runT(t, input, tests)
}

func TestLexer_NextToken_Full(t *testing.T) {
	input := `
		let five = 5;
		let ten = 10;
		
		let add = fn(x, y) {
			x + y;
		};
		let result = add(five, ten);
		!-/*5;
		5 < 10 > 5;

		if (5 < 10) {
			return true;
		} else {
			return false;
		}

		10 == 10;
		10 != 91;
		"foobar"
		"foo bar"
		[1, 2]
		{"foo": "bar"}
	`
	tests := []struct {
		ExpectedType    token.Type
		ExpectedLiteral string
	}{
		{ExpectedType: token.LET, ExpectedLiteral: "let"},
		{ExpectedType: token.IDENT, ExpectedLiteral: "five"},
		{ExpectedType: token.ASSIGN, ExpectedLiteral: "="},
		{ExpectedType: token.INT, ExpectedLiteral: "5"},
		{ExpectedType: token.SEMICOLON, ExpectedLiteral: ";"},

		{ExpectedType: token.LET, ExpectedLiteral: "let"},
		{ExpectedType: token.IDENT, ExpectedLiteral: "ten"},
		{ExpectedType: token.ASSIGN, ExpectedLiteral: "="},
		{ExpectedType: token.INT, ExpectedLiteral: "10"},
		{ExpectedType: token.SEMICOLON, ExpectedLiteral: ";"},

		{ExpectedType: token.LET, ExpectedLiteral: "let"},
		{ExpectedType: token.IDENT, ExpectedLiteral: "add"},
		{ExpectedType: token.ASSIGN, ExpectedLiteral: "="},
		{ExpectedType: token.FUNCTION, ExpectedLiteral: "fn"},
		{ExpectedType: token.LPAREN, ExpectedLiteral: "("},
		{ExpectedType: token.IDENT, ExpectedLiteral: "x"},
		{ExpectedType: token.COMMA, ExpectedLiteral: ","},
		{ExpectedType: token.IDENT, ExpectedLiteral: "y"},
		{ExpectedType: token.RPAREN, ExpectedLiteral: ")"},
		{ExpectedType: token.LBRACE, ExpectedLiteral: "{"},
		{ExpectedType: token.IDENT, ExpectedLiteral: "x"},
		{ExpectedType: token.PLUS, ExpectedLiteral: "+"},
		{ExpectedType: token.IDENT, ExpectedLiteral: "y"},
		{ExpectedType: token.SEMICOLON, ExpectedLiteral: ";"},
		{ExpectedType: token.RBRACE, ExpectedLiteral: "}"},
		{ExpectedType: token.SEMICOLON, ExpectedLiteral: ";"},

		{ExpectedType: token.LET, ExpectedLiteral: "let"},
		{ExpectedType: token.IDENT, ExpectedLiteral: "result"},
		{ExpectedType: token.ASSIGN, ExpectedLiteral: "="},
		{ExpectedType: token.IDENT, ExpectedLiteral: "add"},
		{ExpectedType: token.LPAREN, ExpectedLiteral: "("},
		{ExpectedType: token.IDENT, ExpectedLiteral: "five"},
		{ExpectedType: token.COMMA, ExpectedLiteral: ","},
		{ExpectedType: token.IDENT, ExpectedLiteral: "ten"},
		{ExpectedType: token.RPAREN, ExpectedLiteral: ")"},
		{ExpectedType: token.SEMICOLON, ExpectedLiteral: ";"},

		{ExpectedType: token.BANG, ExpectedLiteral: "!"},
		{ExpectedType: token.MINUS, ExpectedLiteral: "-"},
		{ExpectedType: token.SLASH, ExpectedLiteral: "/"},
		{ExpectedType: token.ASTERISK, ExpectedLiteral: "*"},
		{ExpectedType: token.INT, ExpectedLiteral: "5"},
		{ExpectedType: token.SEMICOLON, ExpectedLiteral: ";"},

		{ExpectedType: token.INT, ExpectedLiteral: "5"},
		{ExpectedType: token.LT, ExpectedLiteral: "<"},
		{ExpectedType: token.INT, ExpectedLiteral: "10"},
		{ExpectedType: token.GT, ExpectedLiteral: ">"},
		{ExpectedType: token.INT, ExpectedLiteral: "5"},
		{ExpectedType: token.SEMICOLON, ExpectedLiteral: ";"},

		{ExpectedType: token.IF, ExpectedLiteral: "if"},
		{ExpectedType: token.LPAREN, ExpectedLiteral: "("},
		{ExpectedType: token.INT, ExpectedLiteral: "5"},
		{ExpectedType: token.LT, ExpectedLiteral: "<"},
		{ExpectedType: token.INT, ExpectedLiteral: "10"},
		{ExpectedType: token.RPAREN, ExpectedLiteral: ")"},
		{ExpectedType: token.LBRACE, ExpectedLiteral: "{"},

		{ExpectedType: token.RETURN, ExpectedLiteral: "return"},
		{ExpectedType: token.TRUE, ExpectedLiteral: "true"},
		{ExpectedType: token.SEMICOLON, ExpectedLiteral: ";"},

		{ExpectedType: token.RBRACE, ExpectedLiteral: "}"},
		{ExpectedType: token.ELSE, ExpectedLiteral: "else"},
		{ExpectedType: token.LBRACE, ExpectedLiteral: "{"},

		{ExpectedType: token.RETURN, ExpectedLiteral: "return"},
		{ExpectedType: token.FALSE, ExpectedLiteral: "false"},
		{ExpectedType: token.SEMICOLON, ExpectedLiteral: ";"},

		{ExpectedType: token.RBRACE, ExpectedLiteral: "}"},

		{ExpectedType: token.INT, ExpectedLiteral: "10"},
		{ExpectedType: token.EQ, ExpectedLiteral: "=="},
		{ExpectedType: token.INT, ExpectedLiteral: "10"},
		{ExpectedType: token.SEMICOLON, ExpectedLiteral: ";"},

		{ExpectedType: token.INT, ExpectedLiteral: "10"},
		{ExpectedType: token.NEQ, ExpectedLiteral: "!="},
		{ExpectedType: token.INT, ExpectedLiteral: "91"},
		{ExpectedType: token.SEMICOLON, ExpectedLiteral: ";"},

		{ExpectedType: token.STRING, ExpectedLiteral: "foobar"},
		{ExpectedType: token.STRING, ExpectedLiteral: "foo bar"},

		{ExpectedType: token.LBRACKET, ExpectedLiteral: "["},
		{ExpectedType: token.INT, ExpectedLiteral: "1"},
		{ExpectedType: token.COMMA, ExpectedLiteral: ","},
		{ExpectedType: token.INT, ExpectedLiteral: "2"},
		{ExpectedType: token.RBRACKET, ExpectedLiteral: "]"},

		{ExpectedType: token.LBRACE, ExpectedLiteral: "{"},
		{ExpectedType: token.STRING, ExpectedLiteral: "foo"},
		{ExpectedType: token.COLON, ExpectedLiteral: ":"},
		{ExpectedType: token.STRING, ExpectedLiteral: "bar"},
		{ExpectedType: token.RBRACE, ExpectedLiteral: "}"},

		{ExpectedType: token.EOF, ExpectedLiteral: ""},
	}
	runT(t, input, tests)
}

func runT(t *testing.T, input string, tests []struct {
	ExpectedType    token.Type
	ExpectedLiteral string
}) {
	lexer := New(input)
	for i, tt := range tests {
		nextToken := lexer.NextToken()
		if nextToken.Type != tt.ExpectedType {
			t.Fatalf("[TOKEN] - Expected %s but got %s at position %d", tt.ExpectedType, nextToken.Type, i)
		}

		if nextToken.Literal != tt.ExpectedLiteral {
			t.Fatalf("[LITERAL] - Expected %s but got %s at position %d", tt.ExpectedLiteral, nextToken.Literal, i)
		}
	}
}
