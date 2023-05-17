package interpreter

import "../ast"

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
	return i.evaluate(grouping.Expr)
}
