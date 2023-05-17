package parser

import (
	"../ast"
	parserror "../parsererror"
	"../token"
)

type Parser struct {
	tokens  []*token.Token
	current int64
}

func (p *Parser) expression() ast.Binary {
	return p.equality()
}

func (p *Parser) equality() ast.Binary {
	expr := p.comparison()

	for p.match(token.BANG, token.BANG_EQUAL) {
		operator := p.previous()
		right := p.comparison()
		expr = ast.Binary{Left: expr, Operator: operator, Right: right}
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

func (p *Parser) advance() token.Token {
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

func (p *Parser) previous() token.Token {
	return *p.tokens[p.current-1]
}

func (p *Parser) comparison() ast.Binary {
	expr := p.term()
	for p.match(token.GREATER, token.GREATER_EQUAL, token.LESS, token.LESS_EQUAL) {
		operator := p.previous()
		right := p.term()
		expr = ast.Binary{Left: expr, Operator: operator, Right: right}
	}
	return expr
}

func (p *Parser) term() ast.Binary {
	expr := p.factor()

	for p.match(token.MINUS, token.PLUS) {
		operator := p.previous()
		right := p.factor()
		expr = ast.Binary{Left: expr, Operator: operator, Right: right}
	}
	return expr
}

func (p *Parser) factor() ast.Binary {
	unary := p.unary()

	var (
		expr ast.Binary
	)

	for p.match(token.SLASH, token.STAR) {
		operator := p.previous()
		right := p.unary()
		expr = ast.Binary{Expr: unary, Operator: operator, Right: right}
	}
	return expr
}

func (p *Parser) unary() ast.Expression {
	if p.match(token.BANG, token.MINUS) {
		operator := p.previous()
		right := p.unary()
		return ast.Unary{
			Operator: operator,
			Right:    right,
		}
	}
	return p.primary()
}

// @TODO: Cleanup this function
func (p *Parser) primary() *ast.Literal {
	if p.match(token.FALSE) {
		return &ast.Literal{Value: false}
	}
	if p.match(token.TRUE) {
		return &ast.Literal{Value: true}
	}
	if p.match(token.NIL) {
		return &ast.Literal{Value: nil}
	}
	if p.match(token.NUMBER, token.STRING) {
		return &ast.Literal{Value: p.previous().Literal}
	}
	if p.match(token.LEFT_PAREN) {
		expr := p.expression()
		p.consume(token.RIGHT_PAREN, "Expect ')' after expression.")
		return &ast.Grouping{Expr: expr}
	}
	return nil
}

func (p *Parser) consume(atype token.TokenType, msg string) (token.Token, error) {
	if p.check(atype) {
		return p.advance(), nil
	}
	return token.Token{}, parserror.MakeError(*p.peek(), msg)
}

func (p *Parser) check(atype token.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().TokenType == atype
}

/* The goal of synchronize is to disgard tokens until we've reached the beginning of the next statement*/
func (p *Parser) synchronize() {
	p.advance()
	for !p.isAtEnd() {
		if p.previous().TokenType == token.SEMICOLON {
			return
		}
		switch p.peek().TokenType {
		case token.CLASS, token.FUN, token.VAR, token.FOR, token.IF, token.WHILE, token.PRINT, token.RETURN:
			return
		}
		p.advance()
	}
}
