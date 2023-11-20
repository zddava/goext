package collection

type Set[T comparable] map[T]bool

func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

func AddToSet[T comparable](set Set[T], value T) Set[T] {
	set[value] = true
	return set
}

func SetValues[T comparable](set Set[T]) (list []T) {
	for k := range set {
		list = append(list, k)
	}

	return list
}
