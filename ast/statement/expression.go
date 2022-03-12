/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/11-11:55 AM
File: expression.go
*/

package statement

import (
	"zlang/ast"
	"zlang/token"
)

// Expression represents a single expression which implement the Stat and Node interface.
type Expression struct {
	Token      token.Token // the first token of the expression
	Expression ast.Expr
}

func (e *Expression) IsStat()         {}
func (e *Expression) Literal() string { return e.Token.Literal }
func (e *Expression) String() string {
	if e.Expression != nil {
		return e.Expression.String()
	}
	return e.Token.Literal
}
