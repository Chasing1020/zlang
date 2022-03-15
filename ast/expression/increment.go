/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/15-1:17 AM
File: increment.go
*/

package expression

import (
	"fmt"
	"zlang/ast"
	"zlang/token"
)

// increment implement the Expr and Assignable interface.
type Increment struct {
	Token token.Token // The [ token
	Left  ast.Expr
	Index ast.Expr
}

func (i *Increment) IsExpr()         {}
func (i *Increment) Literal() string { return i.Token.Literal }
func (i *Increment) String() string {
	return fmt.Sprintf("(%s++))", i.Left.String(), i.Index.String())
}
