package syntax

import (
	"../token"
)

type Node interface {
	String() string
}

type Expr interface {
	Node
}

type Binary struct {
	left     Expr
	right    Expr
	operator token.Token
}
