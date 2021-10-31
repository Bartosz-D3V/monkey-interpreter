package object

import (
	"bytes"
	"monkey_interpreter/ast"
	"strings"
)

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() Type {
	return FunctionObj
}

func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := make([]string, 0)
	for _, param := range f.Parameters {
		params = append(params, param.String())
	}
	out.WriteString("fn")
	out.WriteByte('(')
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}
