package lexer

import (
	"monkey_interpreter/token"
)

type Lexer struct {
	input        string // input source code to parse
	position     int    // position current, parsed position - corresponds to ch value
	readPosition int    // readPosition next position that should be parsed
	ch           byte   // ch value of index position from input string

}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tk token.Token
	l.skipWhitespace()
	ch := l.ch
	switch ch {
	case '=':
		if l.input[l.readPosition] == '=' {
			logicOp := l.readLogicOp()
			tk = token.Token{Type: token.EQ, Literal: logicOp}
		} else {
			tk = token.Token{Type: token.ASSIGN, Literal: string(ch)}
		}
	case '+':
		tk = token.Token{Type: token.PLUS, Literal: string(ch)}
	case '(':
		tk = token.Token{Type: token.LPAREN, Literal: string(ch)}
	case ')':
		tk = token.Token{Type: token.RPAREN, Literal: string(ch)}
	case '{':
		tk = token.Token{Type: token.LBRACE, Literal: string(ch)}
	case '}':
		tk = token.Token{Type: token.RBRACE, Literal: string(ch)}
	case ',':
		tk = token.Token{Type: token.COMMA, Literal: string(ch)}
	case ';':
		tk = token.Token{Type: token.SEMICOLON, Literal: string(ch)}
	case '"':
		value := l.readString()
		tk = token.Token{Type: token.STRING, Literal: value}
	case '!':
		if l.input[l.readPosition] == '=' {
			logicOp := l.readLogicOp()
			tk = token.Token{Type: token.NEQ, Literal: logicOp}
		} else {
			tk = token.Token{Type: token.BANG, Literal: string(ch)}
		}
	case '-':
		tk = token.Token{Type: token.MINUS, Literal: string(ch)}
	case '/':
		tk = token.Token{Type: token.SLASH, Literal: string(ch)}
	case '*':
		tk = token.Token{Type: token.ASTERISK, Literal: string(ch)}
	case '<':
		tk = token.Token{Type: token.LT, Literal: string(ch)}
	case '>':
		tk = token.Token{Type: token.GT, Literal: string(ch)}
	case 0:
		tk = token.Token{Type: token.EOF, Literal: ""}
	default:
		if isChar(ch) {
			ident := l.readIdent()
			identType := lookupIdent(ident)
			tk = token.Token{Type: identType, Literal: ident}
			return tk
		} else if isDigit(ch) {
			num := l.readNum()
			tk = token.Token{Type: token.INT, Literal: num}
			return tk
		} else {
			tk = token.Token{Type: token.ILLEGAL, Literal: "ILLEGAL"}
		}
	}
	l.readChar()
	return tk
}
