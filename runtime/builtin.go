/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/13-4:29 PM
File: builtin.go
*/

package runtime

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	builtinFunctions["type"] = func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("wrong number of arguments. got=%d, want=1", len(args))
		}
		objectType, ok := object.TypeMap[args[0].Type()]
		if !ok {
			objectType = "UNKNOWN"
		}
		return &object.String{Value: objectType}
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

	builtinFunctions["newArray"] = func(args ...object.Object) object.Object {
		if len(args) != 1 {
			newError("wrong number of arguments. got=%d, want=1", len(args))
		}
		length, err := strconv.Atoi(args[0].String())
		if err != nil {
			newError("param can't be negative to a integer")
		}
		elements := make([]object.Object, length)
		for i := 0; i < length; i++ {
			elements[i] = &object.Integer{Value: 0}
		}
		return &object.Array{Elements: elements}
	}

	builtinFunctions["print"] = func(args ...object.Object) object.Object {
		for _, arg := range args {
			fmt.Print(arg.String())
		}
		return &object.Null{}
	}

	builtinFunctions["println"] = func(args ...object.Object) object.Object {
		for _, arg := range args {
			fmt.Println(arg.String())
		}
		return &object.Null{}
	}

	builtinFunctions["input"] = func(args ...object.Object) object.Object {
		if len(args) != 0 {
			newError("input error: wrong number of arguments. got=%d", len(args))
		}
		scanner := bufio.NewScanner(os.Stdin)
		ok := scanner.Scan()
		if !ok {
			newError("input error: except io.EOF failed")
		}
		return &object.String{Value: scanner.Text()}
	}

	builtinFunctions["string"] = func(args ...object.Object) object.Object {
		if len(args) != 1 {
			newError("string cast error: int() only accepts one argument")
		}
		return &object.String{Value: args[0].String()}
	}

	builtinFunctions["int"] = func(args ...object.Object) object.Object {
		if len(args) != 1 {
			newError("int cast error: int() only accepts one argument")
		}
		integer, err := strconv.Atoi(args[0].String())
		if err != nil {
			newError("invalid literal for int() with base 10: %s", args[0].String())
		}
		return &object.Integer{Value: integer}
	}
}
