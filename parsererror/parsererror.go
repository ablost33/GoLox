package parserror

import (
	"fmt"

	"../token"
)

func MakeError(tok token.Token, message string) error {
	if tok.TokenType == token.EOF {
		return fmt.Errorf("[line %v] Error at end: %s", tok.Line, message)
	}
	return fmt.Errorf("[line %v] Error at '%s': %s", tok.Line, tok.Lexeme, message)
}
