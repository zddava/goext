package enum

import (
	"reflect"
	"unsafe"
)

type Enum struct {
	Code string
	Name string
}

var enumMap map[reflect.Type]map[string]*Enum = make(map[reflect.Type]map[string]*Enum)

func storeEnum[T any](code string, enum *T) {
	typ := reflect.TypeOf(*enum)
	if _, ok := enumMap[typ]; !ok {
		enumMap[typ] = make(map[string]*Enum)
	}

	enumMap[typ][code] = (*Enum)(unsafe.Pointer(enum))
}

func InitEnum[T any](code string, name string) *T {
	enum := (*T)(unsafe.Pointer(&Enum{code, name}))
	storeEnum(code, enum)
	return enum
}

func ParseEnum[T any](code string) *T {
	var enum T
	typ := reflect.TypeOf(enum)

	if _, ok := enumMap[typ]; !ok {
		return nil
	}

	if _, ok := enumMap[typ][code]; !ok {
		return nil
	}

	return (*T)(unsafe.Pointer(enumMap[typ][code]))
}
