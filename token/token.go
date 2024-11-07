package token

import (
	"fmt"
)

type TokenType int

const (
	// Single-character tokens.
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	// One or two character tokens.
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// Literals.
	IDENTIFIER
	STRING
	NUMBER

	// Keywords.
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE
	END // Set when next chars indicate the end of and expr/statement

	// End of file
	EOF
)

type Token struct {
	Type    TokenType
	Lex     string
	Literal *literal.Literal
	Line    int
}

func NewToken(tokType TokenType, lexeme string, literal *literal.Literal, line int) *Token {
	return &Token{
		Type:    tokType,
		Lex:     lex,
		Literal: literal,
		Line:    line,
	}
}

func (t *Token) String() string {
	return fmt.Sprintf("%d %s %s", t.Type, t.Lex, t.Literal.String())
}
