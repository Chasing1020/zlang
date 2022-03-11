/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/11-11:52 AM
File: return.go
*/

package statement

import (
	"bytes"
	"zlang/ast"
	"zlang/token"
)

// Return implement the Node and Stat interface
type Return struct {
	Token token.Token
	Value ast.Expr
}

func (rs *Return) IsStat()         {}
func (rs *Return) Literal() string { return rs.Token.Literal }
func (rs *Return) String() string {
	var out bytes.Buffer

	out.WriteString(rs.Literal() + " ")

	if rs.Value != nil {
		out.WriteString(rs.Value.String())
	}

	out.WriteString(";")

	return out.String()
}
