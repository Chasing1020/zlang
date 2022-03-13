/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/3/12-3:36 PM
File: evaluator.go
*/

package runtime

import (
	"fmt"
	"zlang/ast"
	"zlang/ast/expression"
	"zlang/ast/statement"
	"zlang/object"
)

var (
	NullObject  = &object.Null{}
	TrueObject  = &object.Boolean{Value: true}
	FalseObject = &object.Boolean{Value: false}
)

func evalProgram(program *ast.File, env *object.Env) object.Object {
	var result object.Object
	for _, stat := range program.Stats {
		result = Eval(stat, env)
		switch r := result.(type) {
		case *object.ReturnValue:
			return r.Value
		case *object.Error:
			return r
		}
	}
	return result
}

func evalBlockStatement(block *statement.Block, env *object.Env) object.Object {
	var result object.Object
	for _, stat := range block.Statements {
		result = Eval(stat, env)
		if result != nil {
			rt := result.Type()
			if rt == object.RETURN || rt == object.ERROR {
				return result
			}
		}
	}
	return result
}

func toBooleanObject(input bool) *object.Boolean {
	if input {
		return TrueObject
	}
	return FalseObject
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return newError("unknown operator: %s", operator)
	}
}

func evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.INTEGER && right.Type() == object.INTEGER:
		return evalIntegerInfixExpression(operator, left, right)
	case left.Type() == object.STRING && right.Type() == object.STRING:
		return evalStringInfixExpression(operator, left, right)
	case operator == "==":
		return toBooleanObject(left == right)
	case operator == "!=":
		return toBooleanObject(left != right)
	case left.Type() != right.Type():
		return newError("type mismatch: %s",
			operator)
	default:
		return newError("unknown operator: %s",
			operator)
	}
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TrueObject:
		return FalseObject
	case FalseObject:
		return TrueObject
	case NullObject:
		return TrueObject
	default:
		return FalseObject
	}
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER {
		return newError("unknown operator")
	}

	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}
	case "<":
		return toBooleanObject(leftVal < rightVal)
	case ">":
		return toBooleanObject(leftVal > rightVal)
	case "<=":
		return toBooleanObject(leftVal <= rightVal)
	case ">=":
		return toBooleanObject(leftVal >= rightVal)
	case "==":
		return toBooleanObject(leftVal == rightVal)
	case "!=":
		return toBooleanObject(leftVal != rightVal)
	default:
		return newError("unknown operator: %s %d",
			operator, right.Type())
	}
}

func evalStringInfixExpression(operator string, left, right object.Object) object.Object {
	if operator != "+" {
		return newError("unknown operator: %d %s %d",
			left.Type(), operator, right.Type())
	}
	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value
	return &object.String{Value: leftVal + rightVal}
}

func evalIfExpression(i *statement.If, env *object.Env) object.Object {
	condition := Eval(i.Condition, env)
	if isError(condition) {
		return condition
	}
	if isTruthy(condition) {
		return Eval(i.Consequence, env)
	} else if i.Alternative != nil {
		return Eval(i.Alternative, env)
	} else {
		return NullObject
	}
}

func evalForStatement(f *statement.For, env *object.Env) object.Object {
	Eval(f.InitStat, env)
	condition := Eval(f.Condition, env)
	if isError(condition) {
		return condition
	}
	for isTruthy(condition) {
		Eval(f.Body, env)
		Eval(f.UpdateStat, env)
		condition = Eval(f.Condition, env)
		if isError(condition) {
			return condition
		}
	}
	return NullObject
}

func evalIdentifier(node *expression.Identifier, env *object.Env) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}

	if builtin, ok := builtinFunctions[node.Value]; ok {
		return builtin
	}

	return newError("identifier not found: " + node.Value)
}

func isTruthy(obj object.Object) bool {
	switch obj {
	case NullObject:
		return false
	case TrueObject:
		return true
	case FalseObject:
		return false
	default:
		return true
	}
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR
	}
	return false
}

func evalExpressions(exps []ast.Expr, env *object.Env) []object.Object {
	var result []object.Object

	for _, e := range exps {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}

func applyFunction(fn object.Object, args []object.Object) object.Object {
	switch f := fn.(type) {
	case *object.Function:
		extendedEnv := extendFunctionEnv(f, args)
		evaluated := Eval(f.Body, extendedEnv)
		return unwrapReturnValue(evaluated)
		//return evaluated
	case *object.Builtin:
		return f.Func(args...)
	default:
		return newError("not a function: %d", fn.Type())
	}
}

func extendFunctionEnv(fn *object.Function, args []object.Object) *object.Env {
	env := object.NewEnclosedEnv(fn.Env)

	for paramIdx, param := range fn.Parameters {
		env.Set(param.Value, args[paramIdx])
	}

	return env
}

func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}

	return obj
}

func evalIndexExpression(left, index object.Object) object.Object {
	switch {
	case left.Type() == object.ARRAY && index.Type() == object.INTEGER:
		return evalArrayIndexExpression(left, index)
	case left.Type() == object.MAP:
		return evalHashIndexExpression(left, index)
	default:
		return newError("index operator not supported: %d", left.Type())
	}
}

func evalArrayIndexExpression(array, index object.Object) object.Object {
	arrayObject := array.(*object.Array)
	idx := index.(*object.Integer).Value
	max := len(arrayObject.Elements) - 1

	if idx < 0 || idx > max {
		return NullObject
	}

	return arrayObject.Elements[idx]
}

func evalHashLiteral(node *expression.Map, env *object.Env) object.Object {
	m := &object.Map{Pairs: make(map[object.HashCode]object.Pair)}
	for keyNode, valueNode := range node.Pairs {
		key := Eval(keyNode, env)
		if isError(key) {
			return key
		}

		hashKey, ok := key.(object.Comparable)
		if !ok {
			return newError("unusable as hash key: %d", key.Type())
		}

		value := Eval(valueNode, env)
		if isError(value) {
			return value
		}

		hashed := hashKey.HashCode()
		m.Pairs[hashed] = object.Pair{Key: key, Value: value}
	}
	return m
}

func evalHashIndexExpression(hash, index object.Object) object.Object {
	hashObject := hash.(*object.Map)

	key, ok := index.(object.Comparable)
	if !ok {
		return newError("unusable as hash key: %d", index.Type())
	}
	pair, ok := hashObject.Pairs[key.HashCode()]
	if !ok {
		return NullObject
	}
	return pair.Value
}
