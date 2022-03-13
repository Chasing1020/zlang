/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/11-12:52 PM
File: array.go
*/

package expression

import (
	"fmt"
	"strings"
	"zlang/ast"
	"zlang/token"
)

// Array implement the Node and Expr interface.
type Array struct {
	Token    token.Token // the '[' token
	Elements []ast.Expr
}

func (a *Array) IsExpr()         {}
func (a *Array) Literal() string { return a.Token.Literal }
func (a *Array) String() string {
	var elements []string
	for _, element := range a.Elements {
		elements = append(elements, element.String())
	}
	return fmt.Sprintf("[%s]", strings.Join(elements, ", "))
}
