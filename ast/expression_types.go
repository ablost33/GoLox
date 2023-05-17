package ast

import "go/token"

type Expr interface {
	Accept(visitor Visitor)
}

type Binary struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}

// See refactor here: https://github.com/maleksiuk/golox/blob/master/expr/expr.go#L84
type LogicalOperator struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}
