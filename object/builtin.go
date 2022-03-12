/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/12-2:14 PM
File: builtin.go
*/

package object

import "strconv"

// this file provides documentation for predeclared identifiers in zlang.

type Integer struct {
	Value int
}

func (i *Integer) Type() Type         { return INTEGER }
func (i *Integer) String() string     { return strconv.Itoa(i.Value) }
func (i *Integer) HashCode() HashCode { return HashCode{Type: INTEGER, Value: i.Value} }
func (i *Integer) CompareTo(other *Integer) int {
	if i.Value == other.Value {
		return 0
	}
	if i.Value > other.Value {
		return 1
	} else {
		return -1
	}
}


