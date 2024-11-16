package types

type GlistExpr struct {
	Type string
	// Name token.Token
	Data []Expr
}

func NewGlistExpr(data []Expr) Expr {
	return &GlistExpr{
		Type: "GlistExpr",
		// Name: name,
		Data: data,
	}
}

func (v *GlistExpr) Accept(visitor Visitor) (any, error) {
	return visitor.VisitGlistExpr(v)
}

func (v *GlistExpr) GetType() string {
	return v.Type
}
