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
	Value int
}

func (i *Integer) IsExpr()         {}
func (i *Integer) Literal() string { return i.Token.Literal }
func (i *Integer) String() string  { return i.Token.Literal }
