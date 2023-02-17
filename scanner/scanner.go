package scanner

import (
	"../token"
)

type Scanner struct {
	source string
	token  []token.Token
}

func scanTokens()([]token.Token){
	do {
		start = token
		scanToken()
	} while(!isAtEnd())
}
