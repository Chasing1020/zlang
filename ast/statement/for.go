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
	var initStat, condition, updateStat string
	if f.InitStat == nil {
		initStat = ""
	} else {
		initStat = f.InitStat.String()
	}
	if f.Condition == nil {
		condition = ""
	} else {
		condition = f.Condition.String()
	}
	if f.UpdateStat == nil {
		updateStat = ""
	} else {
		updateStat = f.Condition.String()
	}

	s = fmt.Sprintf("for (%s; %s; %s) ", initStat, condition, updateStat)
	if f.Body != nil {
		s += fmt.Sprintf("%s", f.Body.String())
	}
	return
}
