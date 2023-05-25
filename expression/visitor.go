package ast

type Visitor interface {
	VisitBinary(binary *Binary) interface{}
	VisitGrouping(grouping *Grouping) interface{}
	VisitLiteral(literal *Literal) interface{}
	VisitUnary(unary *Unary) interface{}
	VisitVariable(variable *Variable) interface{}
	VisitAssign(assign *Assignment) interface{}
	VisitLogical(logical *Logical) interface{}
	VisitCall(call *Call) interface{}
}

func (bin *Binary) Accept(visitor Visitor) interface{} {
	return visitor.VisitBinary(bin)
}

func (logical *Logical) Accept(visitor Visitor) interface{} {
	return visitor.VisitLogical(logical)
}

func (grouping *Grouping) Accept(visitor Visitor) interface{} {
	return visitor.VisitGrouping(grouping)
}

func (literal *Literal) Accept(visitor Visitor) interface{} {
	return visitor.VisitLiteral(literal)
}

func (unary *Unary) Accept(visitor Visitor) interface{} {
	return visitor.VisitUnary(unary)
}

func (variable *Variable) Accept(visitor Visitor) interface{} {
	return visitor.VisitVariable(variable)
}

func (assign *Assignment) Accept(visitor Visitor) interface{} {
	return visitor.VisitAssign(assign)
}

func (call *Call) Accept(visitor Visitor) interface{} {
	return visitor.VisitCall(call)
}
