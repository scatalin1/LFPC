package token

type TokenType string

const (
	NEMO = "NEMO" // unknown token/character
	EOF  = "EOF"  // end of file

	// Identifiers and literals
	IDENTIFS = "IDENTIFS" //x,y
	INT      = "INT"      //nums

	// Operators
	EQUAL  = "="
	PLUS   = "+"
	DIVIDE = "/"
	MULT   = "*"
	MINUS  = "-"
	BIG    = ">"
	UNO    = "!"
	SMALL  = "<"
	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAR   = "("
	RPAR   = ")"
	LBRACK = "["
	RBRACK = "]"

	// Keywords
	FUN    = "FUNCTION"
	LET    = "LET"
	IF     = "IF"
	ELSE   = "ELSE"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	RETURN = "RETURN"
	NOT    = "NOT"
)

type Token struct {
	Type  TokenType
	Input string
}

var keywords = map[string]TokenType{
	"fn":     FUN,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
	"not":    NOT,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENTIFS
}
