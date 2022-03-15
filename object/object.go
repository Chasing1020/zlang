/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/12-1:35 PM
File: object.go
*/

package object

import (
	"hash/crc32"
)

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

var TypeMap = map[Type]string{NULL: "Null", ERROR: "Error", INTEGER: "Integer", BOOLEAN: "Boolean",
	STRING: "String", RETURN: "Return", FUNCTION: "Function", BUILTIN: "Builtin", ARRAY: "Array", MAP: "Map"}

type Object interface {
	Type() Type
	String() string
}

type MapKey struct {
	Type  Type
	Value uint64
}

type BuiltinFunction func(args ...Object) Object

func (b BuiltinFunction) Literal() string { return "builtin" }
func (b BuiltinFunction) IsStat()         {}

func (b BuiltinFunction) Type() Type     { return FUNCTION }
func (b BuiltinFunction) String() string { return "builtin function" }

func getHashCode(s string) uint32 {
	var v = crc32.ChecksumIEEE([]byte(s))
	return v
}

// Comparable
// ref: https://go.dev/ref/spec#Comparison_operators
type Comparable interface {
	HashCode() HashCode
}

//
//type Sortable interface {
//	CompareTo(object *Object) int
//}
