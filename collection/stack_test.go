package collection

import (
	"testing"

	"github.com/zddava/goext/assert"
)

// TODO more to test

func TestStack(t *testing.T) {
	stack := NewStack[string]()
	PushToStack(stack, "1")
	PushToStack(stack, "2")
	PushToStack(stack, "3")

	assert.Equals(3, StackLen(stack), t)
	assert.False(IsStackEmpty(stack), t)

	e := PopFromStack(stack)
	assert.Equals("3", e, t)
	assert.Equals(2, StackLen(stack), t)
	assert.False(IsStackEmpty(stack), t)

	e = PopFromStack(stack)
	assert.Equals("2", e, t)
	assert.Equals(1, StackLen(stack), t)
	assert.False(IsStackEmpty(stack), t)

	e = PopFromStack(stack)
	assert.Equals("1", e, t)
	assert.Equals(0, StackLen(stack), t)
	assert.True(IsStackEmpty(stack), t)
}
