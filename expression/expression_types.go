package ast

import "go/token"

type Expr interface {
	Accept(visitor Visitor)
}

type Expression struct {
	Expression Expr
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

type Variable struct {
	Name token.Token
}

type Assignment struct {
	Name  token.Token
	Value Expr
}

type Logical struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}

type Call struct {
	Callee     Expr
	Parenthese token.Token
	Arguments  []Expr
}

type Grouping struct {
	Expression Expr
}

type Literal struct {
	Value interface{}
}

type Unary struct {
	Operator token.Token
	Right    Expr
}
