/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/11-12:39 PM
File: block.go
*/

package statement

import (
	"zlang/ast"
	"zlang/token"
)

// Block represents a single expression which implement the Stat and Node interface.
type Block struct {
	Token      token.Token // the { token
	Statements []ast.Stat
}

func (bs *Block) IsStat()         {}
func (bs *Block) Literal() string { return bs.Token.Literal }
func (bs *Block) String() (s string) {
	for _, statement := range bs.Statements {
		s += statement.String()
	}
	return
}
