package object

type Environment struct {
	store map[string]Object
	outer *Environment
}

func NewEnvironment() *Environment {
	s := make(map[string]Object, 0)
	return &Environment{s, nil}
}

func NewEnclosedEnvironment(env *Environment) *Environment {
	return &Environment{
		store: make(map[string]Object),
		outer: env,
	}
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	for !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
