package evaluator

import (
	"monkey_interpreter/ast"
	"monkey_interpreter/object"
)

func booleanToNativeBoolean(val bool) object.Object {
	if val {
		return TRUE
	}
	return FALSE
}

func evalProgram(program *ast.Program) object.Object {
	var result object.Object
	for _, statement := range program.Statements {
		result = Eval(statement)

		if returnVal, ok := result.(*object.ReturnValue); ok {
			return returnVal.Value
		}
	}

	return result
}

func evalBlockStatement(node *ast.BlockStatement) object.Object {
	var result object.Object
	for _, statement := range node.Statements {
		result = Eval(statement)

		if result != nil && result.Type() == object.ReturnValueObj {
			return result
		}
	}
	return result
}

func evalPrefixExpression(operator string, obj object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(obj)
	case "-":
		return evalMinusOperatorExpression(obj)
	default:
		return NULL
	}
}

func evalMinusOperatorExpression(right object.Object) object.Object {
	switch right.(type) {
	case *object.Integer:
		return &object.Integer{Value: -right.(*object.Integer).Value}
	default:
		return NULL
	}
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func evalInfixExpression(left object.Object, operator string, right object.Object) object.Object {
	switch {
	case left.Type() == object.IntegerObj && right.Type() == object.IntegerObj:
		return evalIntegerInfixExpression(left, operator, right)
	case operator == "==":
		return booleanToNativeBoolean(left == right)
	case operator == "!=":
		return booleanToNativeBoolean(left != right)
	default:
		return NULL
	}
}

func evalIntegerInfixExpression(left object.Object, operator string, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value
	switch operator {
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}
	case "<":
		return booleanToNativeBoolean(leftVal < rightVal)
	case ">":
		return booleanToNativeBoolean(leftVal > rightVal)
	case "==":
		return booleanToNativeBoolean(leftVal == rightVal)
	case "!=":
		return booleanToNativeBoolean(leftVal != rightVal)
	default:
		return NULL
	}
}

func evalIfExpression(node *ast.IfExpression) object.Object {
	condition := Eval(node.Condition)
	if isTruthy(condition) {
		return Eval(node.Consequence)
	} else if node.Alternative != nil {
		return Eval(node.Alternative)
	} else {
		return NULL
	}
}

func isTruthy(obj object.Object) bool {
	switch obj {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}
