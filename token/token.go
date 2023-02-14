package token

type Token struct {
	TokenType int
	Lexeme    string
	Literal   interface{}
	Line      int
}
