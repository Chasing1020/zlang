/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/12-3:28 PM
File: heap.go
*/

package object

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
