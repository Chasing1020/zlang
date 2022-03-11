/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/10-10:29 PM
File: nodes.go
*/

package ast

/*
the definition of the expression
expr:
    expr ('+'|'-'|'*'|'/') expr
    | ('-') expr
    | '(' expr ')'
    | literal;

literal:
    IDENTIFIER
    | integer
    | '(' ( expr (',' expr)* )? ')'; // function call(s)
*/

// Pos represents the position of the node
type Pos struct {
	filename  string
	line, col uint32
}


// Node represents the literal node of an expression or a statement.
type Node interface {
	Literal() string
	String() string
	//Pos() Pos
}

type Expr interface {
	Node
	// IsExpr is used to implement the stat interface.
	IsExpr()
}

type Stat interface {
	Node
	// IsStat is used to implement the stat interface.
	IsStat()
}

// File implements the Node interface.
type File struct {
	Stats []Stat
	EOF Pos
}

func (f *File) Literal() string {
	if len(f.Stats) > 0 {
		return f.Stats[0].Literal()
	} else {
		return ""
	}
}

func (f *File) String() (s string) {
	for _, statement := range f.Stats {
		s += statement.String()
	}
	return
}
