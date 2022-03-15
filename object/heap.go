/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/12-3:28 PM
File: heap.go
*/

package object

import "fmt"

func NewEnclosedEnv(outer *Env) *Env {
	env := NewEnv()
	env.outer = outer
	return env
}

func NewEnv() *Env {
	return &Env{store: make(map[string]Object), outer: nil}
}

type Env struct {
	store map[string]Object
	outer *Env
}

func (e *Env) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Env) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

func (e *Env) SetIndex(name string, index, val Object) Object {
	switch k := e.store[name].(type) {
	case *Array:
		// TODO: support 2d Array
		if index.Type() == INTEGER {
			k.Elements[index.(*Integer).Value] = val
		} else {
			return &Error{Message: fmt.Sprintf("type not supported: %d", index.Type())}
		}
	case *Map:
		key, ok := index.(Comparable)
		if !ok {
			return &Error{Message: fmt.Sprintf("unusable as hash key: %d", index.Type())}
		}
		k.Pairs[key.HashCode()] = Pair{Key: index, Value: val}
	}
	return val
}
