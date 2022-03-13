/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/12-2:14 PM
File: builtin.go
*/

package object

import (
	"fmt"
	"hash/crc32"
	"strconv"
	"strings"
	"zlang/ast/expression"
	"zlang/ast/statement"
)

// this file provides documentation for predeclared identifiers in zlang.

type Integer struct {
	Value int
}

func (i *Integer) Type() Type     { return INTEGER }
func (i *Integer) String() string { return strconv.Itoa(i.Value) }
func (i *Integer) HashCode() HashCode {
	return HashCode{Type: INTEGER, Value: uint64(i.Value)}
}

//func (i *Integer) HashCode() HashCode { return HashCode{Type: INTEGER, Value: i.Value} }
//func (i *Integer) CompareTo(other *Integer) int {
//	if i.Value == other.Value {
//		return 0
//	}
//	if i.Value > other.Value {
//		return 1
//	} else {
//		return -1
//	}
//}

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() Type     { return BOOLEAN }
func (b *Boolean) String() string { return strconv.FormatBool(b.Value) }
func (b *Boolean) HashCode() HashCode {
	var value uint64
	if b.Value {
		value = 1
	} else {
		value = 0
	}
	return HashCode{Type: b.Type(), Value: value}
}


type Null struct{}

func (n *Null) Type() Type     { return NULL }
func (n *Null) String() string { return "null" }


type Error struct {
	Message string
}

func (e *Error) Type() Type     { return ERROR }
func (e *Error) String() string { return e.Message }

type Function struct {
	Parameters []*expression.Identifier
	Body       *statement.Block
	Env        *Env
}

func (f *Function) Type() Type { return FUNCTION }
func (f *Function) String() string {
	var params []string
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	return fmt.Sprintf("function(%s) %s", strings.Join(params, ", "), f.Body.String())
}

type String struct {
	Value string
}

func (s *String) Type() Type     { return STRING }
func (s *String) String() string { return s.Value }
func (s *String) HashCode() HashCode {
	checkSum := crc32.ChecksumIEEE([]byte(s.String()))
	return HashCode{Type: s.Type(), Value: uint64(checkSum)}
}

type Builtin struct {
	Func func(args ...Object) Object
}

func (b *Builtin) Type() Type     { return BUILTIN }
func (b *Builtin) String() string { return "builtin function" }

type Array struct {
	Elements []Object
}

func (a *Array) Type() Type { return ARRAY }
func (a *Array) String() string {
	var elements []string
	for _, element := range a.Elements {
		elements = append(elements, element.String())
	}
	return fmt.Sprintf("[%s]", strings.Join(elements, ", "))
}

// Pair encapsulates the key/value pairs
type Pair struct {
	Key   Object
	Value Object
}

// HashCode an identifier of an object
type HashCode struct {
	Type  Type
	Value uint64
}

// Map
// Two interface values are equal if they have identical dynamic types
// and equal dynamic values or if both have value nil.
// 不能采用Object -> Object，因为每次Eval返回都是新的指向Object的指针
type Map struct {
	Pairs map[HashCode]Pair
}

func (m *Map) Type() Type { return MAP }
func (m *Map) String() (s string) {
	s += "{"
	for _, pair := range m.Pairs {
		s += pair.Key.String() + ": " + pair.Value.String() + ", "
	}
	s = s[:len(s)-1] + "}"
	return
}

type ReturnValue struct {
	Value Object
}

func (r *ReturnValue) Type() Type     { return RETURN }
func (r *ReturnValue) String() string { return r.Value.String() }
