package ast

import "../token"

type Node interface {
	String() string
}

type Expr interface {
	Node
}

type Stmt interface {
	Node
}

type Expression struct {
	Stmt
	Expression Expr
}

type Binary struct {
	Expr
	Left     Expr
	Operator token.Token
	Right    Expr
}
