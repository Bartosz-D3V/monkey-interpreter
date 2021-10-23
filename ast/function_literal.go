package ast

import (
	"bytes"
	"monkey_interpreter/token"
	"strings"
)

type FunctionLiteral struct {
	Token      token.Token // The FN token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

func (fl *FunctionLiteral) expressionNode() {}

func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer
	out.WriteString(fl.TokenLiteral())
	out.WriteByte('(')

	var params []string
	for _, param := range fl.Parameters {
		params = append(params, param.Value)
	}
	out.WriteString(strings.Join(params, ","))
	out.WriteByte(')')
	out.WriteByte('{')
	out.WriteString(fl.Body.String())
	out.WriteByte('}')
	return out.String()
}
