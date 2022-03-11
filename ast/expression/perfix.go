/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/11-11:44 AM
File: perfix.go
*/

package expression

import (
	"fmt"
	"zlang/ast"
	"zlang/token"
)

// Prefix implement the Node and Expr interface.
type Prefix struct {
	Token    token.Token // The prefix token, e.g. !
	Operator string
	Right    ast.Expr
}

func (p *Prefix) IsExpr()         {}
func (p *Prefix) Literal() string { return p.Token.Literal }
func (p *Prefix) String() string {
	return fmt.Sprintf("(%s%s)", p.Operator, p.Right.String())
}
