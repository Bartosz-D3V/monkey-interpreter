package evaluator

import (
	"fmt"
	"monkey_interpreter/object"
)

var builtins = map[string]*object.BuiltIn{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) > 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to `len` not supported, got %s", arg.Type())
			}
		},
	},
	"first": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) > 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			switch arg := args[0].(type) {
			case *object.Array:
				if len(arg.Elements) == 0 {
					return NULL
				}
				return &object.ReturnValue{Value: arg.Elements[0]}
			default:
				return newError("argument to `len` not supported, got %s", arg.Type())
			}
		},
	},
	"last": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) > 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			switch arg := args[0].(type) {
			case *object.Array:
				if len(arg.Elements) == 0 {
					return NULL
				}
				return &object.ReturnValue{Value: arg.Elements[len(arg.Elements)-1]}
			default:
				return newError("argument to `len` not supported, got %s", arg.Type())
			}
		},
	},
	"rest": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) > 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			switch arg := args[0].(type) {
			case *object.Array:
				if len(arg.Elements) == 0 {
					return NULL
				}
				arrElems := make([]object.Object, len(arg.Elements)-1)
				copy(arrElems, arg.Elements[1:])
				return &object.Array{Elements: arrElems}
			default:
				return newError("argument to `len` not supported, got %s", arg.Type())
			}
		},
	},
	"push": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			switch arg := args[0].(type) {
			case *object.Array:
				newElem := args[1]
				return &object.Array{Elements: append(arg.Elements, newElem)}
			default:
				return newError("argument to `push` not supported, got %s", arg.Type())
			}
		},
	},
	"puts": {
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}
			return NULL
		},
	},
}
