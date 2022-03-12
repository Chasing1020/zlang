/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/12-1:35 PM
File: object.go
*/

package object

type Type uint

const (
	_ Type = iota
	NULL
	ERROR
	INTEGER
	BOOLEAN
	STRING
	RETURN
	FUNCTION
	BUILTIN
	ARRAY
	MAP
)

type Object interface {
	Type() Type
	String() string
}


type HashCode struct {
	Type  Type
	Value int
}

// Comparable
// ref: https://go.dev/ref/spec#Comparison_operators
type Comparable interface {
	HashCode() HashCode
}

type Sortable interface {
	CompareTo(object *Object) int
}


//eval(source, globals=None, locals=None, /)
//Evaluate the given source in the context of globals and locals.