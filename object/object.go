package object

type Object interface {
	Type() Type
	Inspect() string
}
