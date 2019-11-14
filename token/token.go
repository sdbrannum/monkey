package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"
	
	// identifier & literals
	IDENT = "IDENT" // variable names e.g x, y, foo
	INT = "INT" // values e.g. 1, 2

	// operators
	ASSIGN = "="
	PLUS = "+"

	// delimiters
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET = "LET"

)