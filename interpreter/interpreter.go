package interpreter

import (
	ast "../expression"
	"../token"
)

type environment struct {
	parent    *environment
	variables map[string]interface{}
}

type Interpreter struct {
	env *environment
}

func (i *Interpreter) VisitLiteral(literal *ast.Literal) interface{} {
	return literal.Value
}

func (i *Interpreter) VisitGrouping(grouping *ast.Grouping) interface{} {
	return i.evaluate(grouping.Expression)
}

func (i *Interpreter) evaluate(expression ast.Expr) interface{} {
	return expression.Accept(i)
}

func (i *Interpreter) VisitUnaryExpr(expr *ast.Unary) interface{} {
	right := i.evaluate(expr.Right)

	switch expr.Operator.TokenType {
	case token.BANG:
		return !isTruthy(right)
	case token.MINUS:
		value, _ := right.(float64)
		return -value
	}
	return nil
}

/*
	We're following Ruby's rule: false & nil are falsey and everything else is truthy
*/
func isTruthy(object interface{}) bool {
	if object == nil {
		return false
	}
	if b, ok := object.(bool); ok {
		return b
	}
	return true
}
