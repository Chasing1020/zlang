/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/11-11:48 AM
File: infix.go
*/

package expression

import (
	"fmt"
	"zlang/ast"
	"zlang/token"
)

// Infix implement the Node and Expr interface.
type Infix struct {
	Token    token.Token
	Left     ast.Expr
	Operator string
	Right    ast.Expr
}

func (ie *Infix) IsExpr()         {}
func (ie *Infix) Literal() string { return ie.Token.Literal }
func (ie *Infix) String() string {
	return fmt.Sprintf("(%s %s %s)", ie.Left.String(), ie.Operator, ie.Right.String())
}
