/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/13-1:44 PM
File: assignment.go
*/

package statement

import (
	"fmt"
	"zlang/ast"
	"zlang/token"
)

// Assignment implement the Expr and Stat interface.
type Assignment struct {
	Token token.Token
	Left  ast.Assignable
	Right ast.Expr
}

func (a *Assignment) IsStat()         {}
func (a *Assignment) IsExpr()         {}
func (a *Assignment) Literal() string { return a.Token.Literal }
func (a *Assignment) String() string {
	return fmt.Sprintf("%s = %s;", a.Left.String(), a.Right.String())
}
