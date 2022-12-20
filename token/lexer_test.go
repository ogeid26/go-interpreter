package lexer

import {
	"testing"
	"/token"
}

func TestNextToken(t *testing.T){
	input := `=+(){},;`

	tests := []struct {
		expectedType token.TokenType
		expectedLiteral string
	} {
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{TOKEN.LPAREN, "("},
		{TOKEN.RPAREN, ")"},
		{TOKEN.LBRACE, "{"},
		{TOKEN.RBRACE, "}"},
		{TOKEN.COMMA, ","},
		{TOKEN.SEMICOLON, ";"},
		{TOKEN.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if


	}
}