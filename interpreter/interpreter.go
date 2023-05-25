package interpreter

import ast "../expression"

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
