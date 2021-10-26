package object

import "fmt"

type Error struct {
	Message string
}

func (e *Error) Type() Type {
	return ErrorObj
}

func (e *Error) Inspect() string {
	return fmt.Sprintf("Error: %s", e.Message)
}
