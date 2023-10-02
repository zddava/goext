package collection

import (
	"testing"

	"github.com/zddava/goext/assert"
)

// TODO more to test

func TestAddToSlice(t *testing.T) {
	var strs []string
	AddToSlice(&strs, "1", "2")

	assert.SliceNotEmpty(strs, t)
	assert.SliceEquals([]string{"1", "2"}, strs, t)
}
