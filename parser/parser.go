package parser

import (
	"../token"
	"../tools"
)

type Parser struct {
	tokens  []*token.Token
	current int64
}

func (p *Parser) expression() tools.Expressions {
	return p.equality()
}
