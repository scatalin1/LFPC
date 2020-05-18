package lexer

import (
	"compiler/token"
	"testing"
)

//TestNewToken fun
func TestNewToken(t *testing.T) {
	input := `let five = 5 ; 
	< > - * / 
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENTIFS, "five"},
		{token.EQUAL, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.SMALL, "<"},
		{token.BIG, ">"},
		{token.MINUS, "-"},
		{token.MULT, "*"},
		{token.DIVIDE, "/"},
		{token.EOF, ""},
	}
	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("Test[%d] = type wrong. Expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Input != tt.expectedLiteral {
			t.Fatalf("Test[%d] = literal wrong. Expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Input)
		}

	}
}
