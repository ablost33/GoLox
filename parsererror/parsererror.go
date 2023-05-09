package parserror

import (
	"../token"
)

func NewParserError(token token.Token, msg string) error {
	Loxerror(token, msg)
	return ParseError()
}

func Loxerror(token token.Token, msg string) {
	if token.TokenType == token.LEFT_PAREN {
		errpkg.report(token.Line, " at end", msg)
	}
}
