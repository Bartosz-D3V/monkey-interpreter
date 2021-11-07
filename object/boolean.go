package object

import "fmt"

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() Type {
	return BooleanObj
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

func (b *Boolean) HashKey() HashKey {
	var hash uint64
	if b.Value {
		hash = 1
	} else {
		hash = 0
	}
	hashKey := HashKey{
		Type:  b.Type(),
		Value: hash,
	}
	return hashKey
}
