package object

import "hash/fnv"

type String struct {
	Value string
}

func (s String) Type() Type {
	return StringObj
}

func (s String) Inspect() string {
	return s.Value
}

func (s *String) HashKey() HashKey {
	hash := fnv.New64a()
	hash.Write([]byte(s.Value))
	return HashKey{
		Type:  s.Type(),
		Value: hash.Sum64(),
	}
}
