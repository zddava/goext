package assert

import (
	"errors"
	"testing"
)

func TestAllPass(t *testing.T) {
	Equals(1, 1, t)
	Equals("1", "1", t)
	Equals(1.1, 1.1, t)
	Equals(true, true, t)
	SliceEquals([]int{1, 2}, []int{1, 2}, t)
	SliceEquals([]string{"1", "2"}, []string{"1", "2"}, t)
	MapEquals(map[int]int{1: 1, 2: 2}, map[int]int{1: 1, 2: 2}, t)
	MapEquals(map[string]string{"a": "aa", "b": "bb"}, map[string]string{"a": "aa", "b": "bb"}, t)

	NotEquals(1, 2, t)
	SliceNotEquals([]int{1, 2}, []int{3, 4}, t)
	MapNotEquals(map[int]int{1: 1, 2: 2}, map[int]int{1: 1, 2: 3}, t)

	True(1 == (2-1), t)
	False(1 == 2, t)

	Nil(nil, t)
	Nil((*int)(nil), t)
	NotNil(new([]string), t)

	SliceNotEmpty([]int{1, 2}, t)
	SliceEmpty([]int{}, t)

	MapNotEmpty(map[int]int{1: 1}, t)
	MapEmpty(map[int]int{}, t)

	NoError(nil, t)
	HasError(errors.New("some error"), t)
}

func TestAllFail(t *testing.T) {
	Equals(1, 2, t)
	Equals("1", "2", t)
	Equals(1.1, 1.2, t)
	Equals(true, false, t)
	SliceEquals([]int{1, 2}, []int{1, 3}, t)
	SliceEquals([]string{"1", "2"}, []string{"1", "3"}, t)
	MapEquals(map[int]int{1: 1, 2: 2}, map[int]int{1: 1, 2: 3}, t)
	MapEquals(map[string]string{"a": "aa", "b": "bb"}, map[string]string{"a": "aa", "c": "cc"}, t)

	NotEquals(1, 1, t)
	SliceNotEquals([]int{1, 1}, []int{1, 1}, t)
	MapNotEquals(map[int]int{1: 1, 2: 2}, map[int]int{1: 1, 2: 2}, t)

	True(1 == 2, t)
	False(1 == (2-1), t)

	Nil(new([]int), t)
	NotNil(nil, t)
	NotNil((*int)(nil), t)

	SliceNotEmpty([]int{}, t)
	SliceEmpty([]int{1, 2}, t)

	MapNotEmpty(map[int]int{}, t)
	MapEmpty(map[int]int{1: 1}, t)

	NoError(errors.New("some error"), t)
	HasError(nil, t)
}
