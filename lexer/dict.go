package lexer

import "monkey_interpreter/token"

var keywords = map[string]token.Type{
	"let":    token.LET,
	"fn":     token.FUNCTION,
	"if":     token.IF,
	"else":   token.ELSE,
	"return": token.RETURN,
	"true":   token.TRUE,
	"false":  token.FALSE,
}

func lookupIdent(ident string) token.Type {
	if tk, ok := keywords[ident]; ok {
		return tk
	} else {
		return token.IDENT
	}
}
