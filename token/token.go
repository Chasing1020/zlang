/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/10-1:01 PM
File: token.go
*/

package token

import (
	"fmt"
	"reflect"
)

type Type uint

const (
	UNKNOWN Type = iota // undefined behavior
	EOF          // EOF
	Ident
	Int
	String

	Operator // op
	Assign   // =
	Plus     // +
	Minus    // -
	Bang     // !
	Star     // *
	Slash    // /

	Eql // ==
	Neq // !=
	Lss // <
	Leq // <=
	Gtr // >
	Geq // >=

	Lparen    // (
	Lbrack    // [
	Lbrace    // {
	Rparen    // )
	Rbrack    // ]
	Rbrace    // }
	Comma     // ,
	Semi      // ;
	Colon     // :
	Dot       // .
	DotDotDot // ...

	Function // function
	True
	False
	If     // if
	Else   // else
	Let    // let
	Return // return

	For    // for
	Const  // const TODO: support const
)

func (t *Type) String() string {
	reflect.TypeOf(t)
	return fmt.Sprintf("%d", t)
}

// Map used to debug
var Map = []string{"UNKNOWN", "EOF", "Ident", "Int", "String", "Operator", "Assign",
	"Plus", "Minus", "Bang", "Star", "Slash", "Eql", "Neq", "Lss", "Leq", "Gtr",
	"Geq", "Lparen", "Lbrack", "Lbrace", "Rparen", "Rbrack", "Rbrace", "Comma",
	"Semi", "Colon", "Dot", "DotDotDot", "Function", "True", "False", "If", "Else", "Let", "Return", "For"}

type Token struct {
	Type    Type
	Literal string
}

func (t Token) String() string {
	return fmt.Sprintf("type: %s , literal:%s", Map[t.Type], t.Literal)
}

type LitKind uint8

const (
	IntLit LitKind = iota
	StringLit
	Keyword
)

var KeywordMap = map[string]Type{
	"function": Function,
	"let":      Let,
	"true":     True,
	"false":    False,
	"if":       If,
	"else":     Else,
	"return":   Return,
	"for":      For,
}
