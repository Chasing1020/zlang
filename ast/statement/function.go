/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/11-12:46 PM
File: function.go
*/

package statement

import (
	"fmt"
	"strings"
	"zlang/ast/expression"
	"zlang/token"
)

// Function represents the function definition.
// implement the Node and Expr interface.
type Function struct {
	Token      token.Token // The 'fn' token
	Parameters []*expression.Identifier
	Body       *Block
}

// IsExpr represents Function is an expression
// Aim to avoid the import cycle problem.
func (f *Function) IsExpr()         {}

// func (f *Function) IsStat()         {}

func (f *Function) Literal() string { return f.Token.Literal }
func (f *Function) String() string {
	var params []string
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	return fmt.Sprintf("function(%s) %s", strings.Join(params, ", "), f.Body.String())
}
