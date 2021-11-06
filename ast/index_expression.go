package ast

import (
	"bytes"
	"monkey_interpreter/token"
)

type IndexExpression struct {
	Token token.Token
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *IndexExpression) expressionNode() {}

func (ie *IndexExpression) String() string {
	var out bytes.Buffer
	out.WriteByte('(')
	out.WriteString(ie.Left.String())
	out.WriteByte('[')
	out.WriteString(ie.Index.String())
	out.WriteByte(']')
	out.WriteByte(')')
	return out.String()
}
