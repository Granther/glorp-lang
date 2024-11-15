package types

import "glorp/literal"

type LiteralExpr struct {
	Type string
	Val *literal.Literal
}

func NewLiteralExpr(val *literal.Literal) Expr {
	return &LiteralExpr{
		Type: "LiteralExpr",
		Val: val,
	}
}

func (l *LiteralExpr) Accept(visitor Visitor) (any, error) {
	return visitor.VisitLiteralExpr(l)
}

func (v *LiteralExpr) GetType() string {
	return v.Type
}