/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/10-10:29 PM
File: common.go
*/

package ast

import "bytes"

// The base Node interface
type Node interface {
	Literal() string
	String() string
}

// All statement nodes implement this
type Statement interface {
	Node
	statementNode()
}

// All expression nodes implement this
type Expression interface {
	Node
	expressionNode()
}

type File struct {
	Statements []Statement
}

func (f *File) TokenLiteral() string {
	if len(f.Statements) > 0 {
		return f.Statements[0].Literal()
	} else {
		return ""
	}
}

func (f *File) String() string {
	var out bytes.Buffer

	for _, s := range f.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}
