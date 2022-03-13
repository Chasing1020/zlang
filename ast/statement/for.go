/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/13-12:36 PM
File: for.go
*/

package statement

import (
	"fmt"
	"zlang/ast"
	"zlang/token"
)

// For implements the Expr and Node interface
type For struct {
	Token      token.Token
	InitStat   ast.Stat
	Condition  ast.Expr
	Body       *Block
	UpdateStat ast.Stat
}

func (f *For) IsStat()         {}
func (f *For) Literal() string { return f.Token.Literal }
func (f *For) String() (s string) {
	s = fmt.Sprintf("for (%s; %s; %s) ", f.InitStat.String(), f.Condition.String(), f.UpdateStat)
	if f.Body != nil {
		s += fmt.Sprintf("%s", f.Body.String())
	}
	return
}
