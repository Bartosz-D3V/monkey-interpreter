package token

type Type string

type Token struct {
	Type    Type
	Literal string
}

const (
	// ILLEGAL Unknown token
	ILLEGAL = "ILLEGAL"
	// EOF End of file - final, ending token
	EOF = "EOF"

	/*
		Identifiers and Literals
	*/
	IDENT = "IDENT" // foo, bar, x, y, z
	INT   = "INT"   // 123, 5

	/*
		SEPARATORS
	*/
	COMMA     = ","
	SEMICOLON = ";"

	/*
		BRACKETS
	*/
	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	/*
		OPS
	*/
	BANG     = "!"
	PLUS     = "+"
	MINUS    = "-"
	ASSIGN   = "="
	SLASH    = "/"
	ASTERISK = "*"
	LT       = "<"
	GT       = ">"

	/*
		KEYWORDS
	*/
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	STRING   = "STRING"

	/*
		LOGIC OPS
	*/
	EQ  = "=="
	NEQ = "!="
)
