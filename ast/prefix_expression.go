package ast

import (
	"bytes"
	"monkey_interpreter/token"
)

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) expressionNode() {}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteByte('(')
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteByte(')')
	return out.String()
}
