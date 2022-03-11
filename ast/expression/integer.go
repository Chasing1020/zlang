/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/11-12:46 AM
File: integer.go
*/

package expression

import "zlang/token"

// Integer implement the Node and Expr interface
type Integer struct {
	Token token.Token
	Value int64
}

func (il *Integer) IsExpr()         {}
func (il *Integer) Literal() string { return il.Token.Literal }
func (il *Integer) String() string  { return il.Token.Literal }
