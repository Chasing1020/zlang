/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/11-11:50 AM
File: identifier.go
*/

package expression

import "zlang/token"

// Identifier implement the Node and Expr interface
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) IsExpr()         {}
func (i *Identifier) IsAssignable()   {}
func (i *Identifier) Literal() string { return i.Token.Literal }
func (i *Identifier) String() string  { return i.Value }
