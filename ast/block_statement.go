package ast

import (
	"bytes"
	"monkey_interpreter/token"
)

type BlockStatement struct {
	Token      token.Token // the '{' token
	Statements []Statement
}

func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}

func (bs *BlockStatement) statementNode() {}

func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for _, statement := range bs.Statements {
		out.WriteString(statement.String())
	}
	return out.String()
}
