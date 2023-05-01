package parser

import (
	"../token"
)

type Parser struct {
	tokens  []*token.Token
	current int64
}

func (p *Parser) expression() ast.Expression {
	return p.equality()
}

func (p *Parser) equality() ast.Expression {
	expr := p.comparison()

	for p.match(token.BANG, token.BANG_EQUAL) {
		operator := p.previous()
		right := p.comparison()
		expr = createNewBinary(expr, operator, right)
	}

	return expr
}

func (p *Parser) match(types ...token.TokenType) bool {
	for _, aType := range types {
		if p.typeIsDefined(aType) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) typeIsDefined(aType token.TokenType) bool {
	if p.isAtEnd() {
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

func (p *Parser) previous() *token.Token {
	return p.tokens[p.current-1]
}

func (p *Parser) comparison() *ast.Expression {
	expr := p.term()
	for p.match(token.GREATER, token.GREATER_EQUAL, token.LESS, token.LESS_EQUAL) {
		operator := p.previous()
		right := p.term()
		expr = createNewBinary(expr, operator, right)
	}
	return expr
}

func (p *Parser) term() *ast.Expression {
	expr := p.factor()

	for p.match(token.MINUS, token.PLUS) {
		operator := p.previous()
		right := p.factor()
		expr = createNewBinary(expr, operator, right)
	}
	return expr
}

func (p *Parser) factor() *ast.Expression {
	expr := p.unary()

	for p.match(token.SLASH, token.STAR) {
		operator := p.previous()
		right := p.unary()
		expr = createNewBinary(expr, operator, right)
	}
	return expr
}

func (p *Parser) unary() *ast.Expression {

}
