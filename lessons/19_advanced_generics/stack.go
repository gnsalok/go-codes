package advancedgenerics

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(value T) {
	s.items = append(s.items, value)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	last := len(s.items) - 1
	value := s.items[last]
	s.items = s.items[:last]
	return value, true
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}
