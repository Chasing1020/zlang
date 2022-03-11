/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/11-12:54 PM
File: index.go
*/

package expression

import (
	"fmt"
	"zlang/ast"
	"zlang/token"
)

// Index implement the Node and Expr interface.
type Index struct {
	Token token.Token // The [ token
	Left  ast.Expr
	Index ast.Expr
}

func (i *Index) IsExpr()         {}
func (i *Index) Literal() string { return i.Token.Literal }
func (i *Index) String() string {
	return fmt.Sprintf("(%s[%s])", i.Left.String(), i.Index.String())
}
