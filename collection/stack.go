package collection

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0),
	}
}

func PushToStack[T any](s *Stack[T], e T) {
	s.data = append(s.data, e)
}

func PopFromStack[T any](s *Stack[T]) T {
	ret := s.data[StackLen(s)-1]
	s.data = s.data[0 : StackLen(s)-1]
	return ret
}

func PeekFromStack[T any](s *Stack[T]) T {
	return s.data[StackLen(s)-1]
}

func StackLen[T any](s *Stack[T]) int {
	return len(s.data)
}

func IsStackEmpty[T any](s *Stack[T]) bool {
	return StackLen(s) == 0
}
