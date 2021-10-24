package ast

import (
	"bytes"
	"monkey_interpreter/token"
	"strings"
)

type CallExpression struct {
	Token     token.Token // The '(' token
	Function  Expression  // The name of the function
	Arguments []Expression
}

func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}

func (ce *CallExpression) expressionNode() {}

func (ce *CallExpression) String() string {
	var out bytes.Buffer
	var args []string
	for _, arg := range ce.Arguments {
		args = append(args, arg.String())
	}
	out.WriteString(ce.Function.TokenLiteral())
	out.WriteByte('(')
	out.WriteString(strings.Join(args, ", "))
	out.WriteByte(')')
	return out.String()
}
