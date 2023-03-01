package token

type Token struct {
	TokenType TokenType
	Lexeme    string
	Literal   interface{}
	Line      int
}
