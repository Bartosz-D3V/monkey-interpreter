package ast

import (
	"bytes"
	"monkey_interpreter/token"
	"strings"
)

type ArrayLiteral struct {
	Token    token.Token
	Elements []Expression
}

func (al *ArrayLiteral) TokenLiteral() string {
	return al.Token.Literal
}

func (al *ArrayLiteral) expressionNode() {}

func (al *ArrayLiteral) String() string {
	var out bytes.Buffer
	var elems []string
	for _, elem := range al.Elements {
		elems = append(elems, elem.String())
	}
	out.WriteByte('[')
	out.WriteString(strings.Join(elems, ", "))
	out.WriteByte(']')
	return out.String()
}
