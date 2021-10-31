package object

type String struct {
	Value string
}

func (s String) Type() Type {
	return StringObj
}

func (s String) Inspect() string {
	return s.Value
}
