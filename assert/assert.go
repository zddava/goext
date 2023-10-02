package assert

import (
	"reflect"
	"testing"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func Equals[T comparable](expect T, actual T, t *testing.T) {
	if expect != actual {
		t.Errorf("expect %v, got: %v", expect, actual)
		return
	}

	t.Logf("expect %v, got: %v", expect, actual)
}

func SliceEquals[T comparable](expect []T, actual []T, t *testing.T) {
	if !slices.Equal(expect, actual) {
		t.Errorf("expect %v, got: %v", expect, actual)
		return
	}

	t.Logf("expect %v, got: %v", expect, actual)
}

func MapEquals[K, V comparable](expect map[K]V, actual map[K]V, t *testing.T) {
	if !maps.Equal(expect, actual) {
		t.Errorf("expect %v, got: %v", expect, actual)
		return
	}

	t.Logf("expect %v, got: %v", expect, actual)
}

func NotEquals[T comparable](expect T, actual T, t *testing.T) {
	if expect == actual {
		t.Errorf("expect not %v, got: %v", expect, actual)
		return
	}

	t.Logf("expect not %v, got: %v", expect, actual)
}

func SliceNotEquals[T comparable](expect []T, actual []T, t *testing.T) {
	if slices.Equal(expect, actual) {
		t.Errorf("expect not %v, got: %v", expect, actual)
		return
	}

	t.Logf("expect not %v, got: %v", expect, actual)
}

func MapNotEquals[K, V comparable](expect map[K]V, actual map[K]V, t *testing.T) {
	if maps.Equal(expect, actual) {
		t.Errorf("expect not %v, got: %v", expect, actual)
		return
	}

	t.Logf("expect not %v, got: %v", expect, actual)
}

func True(actual bool, t *testing.T) {
	if !actual {
		t.Errorf("expect true, got: %v", actual)
	}
}

func False(actual bool, t *testing.T) {
	if actual {
		t.Errorf("expect false, got: %v", actual)
	}
}

func Nil(actual any, t *testing.T) {
	if actual == nil {
		t.Logf("expect nil, got: %v", actual)
		return
	}

	val := reflect.ValueOf(actual)
	if !val.IsNil() {
		t.Errorf("expect nil, got: %v", actual)
	}
}

func NotNil(actual any, t *testing.T) {
	if actual == nil {
		t.Errorf("expect not nil, got: nil")
		return
	}

	val := reflect.ValueOf(actual)
	if val.IsNil() {
		t.Errorf("expect not nil, got: nil")
		return
	}

	t.Logf("expect not nil, got: %v", actual)
}

func SliceNotEmpty[T any](actual []T, t *testing.T) {
	if len(actual) == 0 {
		t.Errorf("expect not empty slice")
	}
}

func MapNotEmpty[K comparable, V any](actual map[K]V, t *testing.T) {
	if len(actual) == 0 {
		t.Errorf("expect not empty map")
	}
}

func SliceEmpty[T any](actual []T, t *testing.T) {
	if len(actual) != 0 {
		t.Errorf("expect empty slice, got: %v", actual)
	}
}

func MapEmpty[K comparable, V any](actual map[K]V, t *testing.T) {
	if len(actual) != 0 {
		t.Errorf("expect empty map, got: %v", actual)
	}
}

func NoError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("expect no error, got: %v", err)
	}
}

func HasError(err error, t *testing.T) {
	if err == nil {
		t.Errorf("expect error, got: nil")
	}

	t.Logf("LOG: expect error, got: %v", err)
}
