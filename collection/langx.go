package collection

func AddToSlice[T any](slice *[]T, value ...T) {
	if slice == nil {
		*slice = make([]T, 0)
	}

	*slice = append(*slice, value...)
}

func EqualsAny[T comparable](one T, values ...T) bool {
	for _, v := range values {
		if one == v {
			return true
		}
	}
	return false
}
