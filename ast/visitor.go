package ast

type Visitor interface {
	VisitBinary(binary *Binary) interface{}
	VisitGrouping(grouping *Grouping) interface{}
	VisitLiteral(literal *Literal) interface{}
	VisitUnary(unary *Unary) interface{}
	VisitVariable(variable *Variable) interface{}
	VisitAssign(assign *Assign) interface{}
	VisitLogical(logical *Logical) interface{}
	VisitCall(call *Call) interface{}
}

func (bin *Binary) Accept(visitor Visitor) interface{} {
	return visitor.VisitBinary(bin)
}
