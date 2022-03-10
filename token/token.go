package token

import (
	"fmt"
	"reflect"
)

type Token uint

const (
	UNKNOWN Token = iota
	EOF           // EOF
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

	Const // const TODO: support const
)

var TokenMap = []string{"UNKNOWN", "EOF", "Ident", "Int", "String", "Operator", "Assign",
	"Plus", "Minus", "Bang", "Star", "Slash", "Eql", "Neq", "Lss", "Leq", "Gtr",
	"Geq", "Lparen", "Lbrack", "Lbrace", "Rparen", "Rbrack", "Rbrace", "Comma",
	"Semi", "Colon", "Dot", "DotDotDot", "Function", "True", "False", "If", "Else", "Let", "Return"}

func (t *Token) String() string {
	reflect.TypeOf(t)
	return fmt.Sprintf("%d", t)
}

type LitKind uint8

const (
	IntLit LitKind = iota
	StringLit
	Keyword
)

var KeywordMap = map[string]Token{
	"function": Function,
	"let":      Let,
	"true":     True,
	"false":    False,
	"if":       If,
	"else":     Else,
	"return":   Return,
}
