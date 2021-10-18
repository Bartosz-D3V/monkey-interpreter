package ast

import (
	"github.com/stretchr/testify/assert"
	"monkey_interpreter/token"
	"testing"
)

func TestProgram_String(t *testing.T) {
	exp := "let myVar = anotherVar;"
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{
					Type:    token.LET,
					Literal: "let",
				},
				Name: &Identifier{
					Token: token.Token{
						Type:    token.IDENT,
						Literal: "myVar",
					},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{
						Type:    token.IDENT,
						Literal: "anotherVar",
					},
					Value: "anotherVar",
				},
			},
		},
	}
	assert.Equal(t, exp, program.String())
}
