/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/11-1:00 PM
File: let.go
*/

package statement

import (
	"fmt"
	"zlang/ast"
	"zlang/ast/expression"
	"zlang/token"
)

// Let implement the Node and Stat interface.
type Let struct {
	Token token.Token // the token.LET token
	Name  *expression.Identifier
	Value ast.Expr
}

func (l *Let) IsStat()         {}
func (l *Let) Literal() string { return l.Token.Literal }
func (l *Let) String() string {
	return fmt.Sprintf("%s %s = %s;", l.Literal(), l.Name.String(), l.Value.String())
}
