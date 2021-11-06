package object

import (
	"bytes"
	"strings"
)

type Array struct {
	Elements []Object
}

func (a *Array) Type() Type {
	return ArrayObj
}

func (a *Array) Inspect() string {
	var out bytes.Buffer

	var elems []string
	for _, elem := range a.Elements {
		elems = append(elems, elem.Inspect())
	}

	out.WriteByte('[')
	out.WriteString(strings.Join(elems, ", "))
	out.WriteByte(']')

	return out.String()
}
