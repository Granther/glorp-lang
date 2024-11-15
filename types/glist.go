package types

// import "glorp/token"

// type GlistExpr struct {
// 	Type string
// 	Name token.Token
// 	Data []Expr
// }

// func NewGlistExpr(name token.Token) Expr {
// 	return &GlistExpr{
// 		Type: "GlistExpr",
// 		Name: name,
// 		Data: []Expr{},
// 	}
// }

// func (v *GlistExpr) Accept(visitor Visitor) (any, error) {
// 	return visitor.VisitGlistExpr(v)
// }

// func (v *GlistExpr) GetType() string {
// 	return v.Type
// }