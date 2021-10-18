package ast

import (
	"bytes"
	"fmt"
	"monkey_interpreter/token"
)

type InfixExpression struct {
	Token      token.Token // Operator (i.e. +, -, / etc)
	LeftValue  Expression
	Operator   string
	RightValue Expression
}

func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *InfixExpression) expressionNode() {}

func (ie *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteByte('(')
	out.WriteString(ie.LeftValue.String())
	out.WriteString(fmt.Sprintf(" %s ", ie.Operator))
	out.WriteString(ie.RightValue.String())
	out.WriteByte(')')
	return out.String()
}
