package parser

import (
	"fmt"

	"../token"
)

func MakeError(tok token.Token, msg string) error {
	if tok.TokenType == token.EOF {
		return fmt.Errorf("[line %v] Error at end: %s", tok.Line, msg)
	}
	return fmt.Errorf("[line %v] Error at %s: %s", tok.Line, tok.Lexeme, msg)
}
