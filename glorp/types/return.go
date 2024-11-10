package types

import "glorp/token"

type ReturnExpr struct {
	Keyword token.Token
	Val     Expr
}

func NewReturnExpr(keyword token.Token, val Expr) Expr {
	return &ReturnExpr{
		Keyword: keyword,
		Val: val,
	}
}

func (r *ReturnExpr) Accept(visitor Visitor) (any, error) {
	return visitor.VisitReturnExpr(r)
}
