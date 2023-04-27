package parser

import (
	"../token"
)

type Parser struct {
	tokens  []*token.Token
	current int64
}

func (p *Parser) expression() *Expression {
	return p.equality()
}

func (p *Parser) equality() *Expression {
	expr := p.comparison()

	for p.match(token.BANG, token.BANG_EQUAL) {
		operator := p.previous()
		right := p.comparison()
		expr = createNewBinary(expr, operator, right)
	}

	return expr
}

func (p *Parser) match(types []*token.TokenType) bool {
	for _, aType := range types {
		if typeIsDefined(aType) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) typeIsDefined(aType *token.TokenType) bool {
	if p.IsAtEnd() {
		return false
	}
	return p.peek().TokenType == aType
}

func (p *Parser) advance() *token.Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().TokenType == token.EOF
}

func (p *Parser) peek() *token.Token {
	return p.tokens[p.current]
}

func (p *Parser) previous()
