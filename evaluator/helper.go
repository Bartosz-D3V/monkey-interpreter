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

func evalStatements(statements []ast.Statement) object.Object {
	var result object.Object
	for _, statement := range statements {
		result = Eval(statement)
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
