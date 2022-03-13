/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/12-2:33 PM
File: runtime.go
*/

package runtime

import (
	"errors"
	"zlang/ast"
	"zlang/ast/expression"
	"zlang/ast/statement"
	"zlang/object"
	"zlang/parser"
)

// Eval will translate the given node into an object.Object.
// It's same as the eval() function in some dynamic languages:
// Python: eval(source, globals=None, locals=None, /)
// JavaScript: declare function eval(x: string): any;
func Eval(node ast.Node, env *object.Env) object.Object {
	switch n := node.(type) {
	// Statements
	case *ast.File:
		return evalProgram(n, env)

	case *statement.Block:
		return evalBlockStatement(n, env)

	case *statement.Expression:
		return Eval(n.Expression, env)

	case *statement.Return:
		val := Eval(n.Value, env)
		if isError(val) {
			return val
		}
		return &object.ReturnValue{Value: val}

	case *statement.Let:
		val := Eval(n.Value, env)
		if isError(val) {
			return val
		}
		env.Set(n.Name.Value, val)

	// Expressions
	case *expression.Integer:
		return &object.Integer{Value: n.Value}

	case *expression.String:
		return &object.String{Value: n.Value}

	case *expression.Boolean:
		return toBooleanObject(n.Value)

	case *expression.Prefix:
		right := Eval(n.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(n.Operator, right)

	case *expression.Infix:
		left := Eval(n.Left, env)
		if isError(left) {
			return left
		}

		right := Eval(n.Right, env)
		if isError(right) {
			return right
		}

		return evalInfixExpression(n.Operator, left, right)

	case *statement.If:
		return evalIfExpression(n, env)

	case *statement.For:
		return evalForStatement(n, env)

	case *expression.Identifier:
		return evalIdentifier(n, env)

	case *statement.Function:
		params := n.Parameters
		body := n.Body
		return &object.Function{Parameters: params, Env: env, Body: body}

	case *expression.Call:
		function := Eval(n.Function, env)
		if isError(function) {
			return function
		}

		args := evalExpressions(n.Arguments, env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}

		return applyFunction(function, args)

	case *expression.Array:
		elements := evalExpressions(n.Elements, env)
		if len(elements) == 1 && isError(elements[0]) {
			return elements[0]
		}
		return &object.Array{Elements: elements}

	case *expression.Index:
		left := Eval(n.Left, env)
		if isError(left) {
			return left
		}
		index := Eval(n.Index, env)
		if isError(index) {
			return index
		}
		return evalIndexExpression(left, index)

	case *expression.Map:
		return evalHashLiteral(n, env)

	case *statement.Assignment:
		right := Eval(n.Right, env)
		if isError(right) {
			return right
		}
		_, ok := env.Get(n.Left.String())
		if !ok {
			return &object.Error{Message: "use of undeclared identifier:" + n.Left.String()}
		}
		env.Set(n.Left.String(), right)
	}
	return nil
}

func Run(buf string) error {
	p := parser.Parser{}
	p.Init(buf)
	file := p.ParseFile()
	env := object.NewEnv()
	evaluated := Eval(file, env)
	_, ok := evaluated.(*object.Error)
	if ok {
		return errors.New(evaluated.String())
	}
	return nil
}
