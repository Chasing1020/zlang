/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/11-12:51 PM
File: string.go
*/

package expression

import "zlang/token"

// String implement the Node and Expr interface.
type String struct {
	Token token.Token
	Value string
}

func (s *String) IsExpr()         {}
func (s *String) Literal() string { return s.Token.Literal }
func (s *String) String() string  { return s.Token.Literal }
