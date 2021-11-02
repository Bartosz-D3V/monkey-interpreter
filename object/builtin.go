package object

type BuiltInFunction func(args ...Object) Object

type BuiltIn struct {
	Fn BuiltInFunction
}

func (bi *BuiltIn) Type() Type {
	return BuiltInObj
}

func (bi *BuiltIn) Inspect() string {
	return "builtin function"
}
