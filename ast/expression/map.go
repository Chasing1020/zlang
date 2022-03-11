/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/11-12:57 PM
File: map.go
*/

package expression

import (
	"fmt"
	"strings"
	"zlang/ast"
	"zlang/token"
)

type Hash struct {
	Token token.Token // the '{' token
	Pairs map[ast.Expr]ast.Expr
}

func (h *Hash) expressionNode()      {}
func (h *Hash) TokenLiteral() string { return h.Token.Literal }
func (h *Hash) String() string {
	var pairs []string
	for key, value := range h.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}
	return fmt.Sprintf("{%s}", strings.Join(pairs, ", "))
}
