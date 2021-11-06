package ast

import (
	"bytes"
	"fmt"
	"monkey_interpreter/token"
	"strings"
)

type HashLiteral struct {
	Token token.Token
	Pairs map[Expression]Expression
}

func (hl *HashLiteral) TokenLiteral() string {
	return hl.Token.Literal
}

func (hl *HashLiteral) expressionNode() {}

func (hl *HashLiteral) String() string {
	var out bytes.Buffer
	out.WriteByte('{')

	var pairs []string
	for key, val := range hl.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s : %s", key.String(), val.String()))
	}

	out.WriteString(strings.Join(pairs, ", "))
	out.WriteByte('}')

	return out.String()
}
