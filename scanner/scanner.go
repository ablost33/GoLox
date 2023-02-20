package scanner

import (
	"../token"
)

type Scanner struct {
	source  string
	start   int
	current int
	line    int
	token   []token.Token
}

func (scanner *Scanner) scanTokens() []token.Token {
	for !scanner.isAtEnd() {
		scanner.start = scanner.current
		scanToken()
	}
	scanner.token = append(scanner.token, token.Token{
		TokenType: token.EOF,
		Line:      scanner.line,
	},
	)
	return scanner.token
}

func (scanner *Scanner) isAtEnd() bool {
	return scanner.current >= len(scanner.source)
}
