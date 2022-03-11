/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/11-12:49 PM
File: call.go
*/

package expression

import (
	"fmt"
	"strings"
	"zlang/ast"
	"zlang/token"
)

// Call represents a function call.
// implement the Node and Expr interface.
type Call struct {
	Token     token.Token // The '(' token
	Function  ast.Expr    // Identifier or FunctionLiteral
	Arguments []ast.Expr
}

func (c *Call) IsExpr()         {}
func (c *Call) Literal() string { return c.Token.Literal }
func (c *Call) String() string {
	var args []string
	for _, a := range c.Arguments {
		args = append(args, a.String())
	}
	return fmt.Sprintf("%s(%s)", c.Function.String(), strings.Join(args, ", "))
}
