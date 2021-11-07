package object

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashMethod(t *testing.T) {
	val1 := &String{Value: "foo"}
	diff1 := &String{Value: "foo"}

	val2 := &String{Value: "bar"}
	diff2 := &String{Value: "bar"}

	assert.Equal(t, val1.HashKey(), diff1.HashKey())
	assert.Equal(t, val2.HashKey(), diff2.HashKey())
	assert.NotEqual(t, val1.HashKey(), val2.HashKey())
	assert.NotEqual(t, diff1.HashKey(), diff2.HashKey())
}
