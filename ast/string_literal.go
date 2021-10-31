package ast

import "monkey_interpreter/token"

type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) TokenLiteral() string {
	return sl.Token.Literal
}

func (sl *StringLiteral) expressionNode() {}

func (sl *StringLiteral) String() string {
	return sl.Token.Literal
}
