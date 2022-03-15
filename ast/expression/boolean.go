/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/11-11:50 AM
File: boolean.go
*/

package expression

import "zlang/token"

// Boolean implement the Node and Expr interface
type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) IsExpr()         {}
func (b *Boolean) Literal() string { return b.Token.Literal }
func (b *Boolean) String() string  { return b.Token.Literal }
