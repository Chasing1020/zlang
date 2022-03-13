/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/13-4:29 PM
File: builtin.go
*/

package runtime

import (
	"zlang/object"
	"zlang/parser"
)

// builtinFunctions: keyword -> builtin function
var builtinFunctions = make(map[string]object.BuiltinFunction)

func init() {
	builtinFunctions["eval"] = func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("wrong number of arguments. got=%d, want=1", len(args))
		}
		p := parser.Parser{}
		p.Init(args[0].String())
		file := p.ParseFile()
		env := object.NewEnv()
		return Eval(file, env)
	}

	builtinFunctions["len"] = func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("wrong number of arguments. got=%d, want=1", len(args))
		}
		switch arg := args[0].(type) {
		case *object.Array:
			return &object.Integer{Value: len(arg.Elements)}
		case *object.String:
			return &object.Integer{Value: len(arg.Value)}
		default:
			return newError("argument to `len` not supported, got %d", args[0].Type())
		}
	}
}
