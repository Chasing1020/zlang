/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/11-12:42 PM
File: if.go
*/

package statement

import (
	"fmt"
	"zlang/ast"
	"zlang/token"
)

// If implements the Expr and Node interface
type If struct {
	Token       token.Token
	Condition   ast.Expr
	Consequence *Block
	Alternative *Block
}

// IsExpr represents If is an expression
// Aim to avoid the import cycle problem.
func (i *If) IsExpr()         {}
func (i *If) Literal() string { return i.Token.Literal }
func (i *If) String() (s string) {
	s = fmt.Sprintf("if (%s) %s", i.Condition.String(), i.Consequence.String())
	if i.Alternative != nil {
		s += fmt.Sprintf("else %s", i.Alternative.String())
	}
	return
}
