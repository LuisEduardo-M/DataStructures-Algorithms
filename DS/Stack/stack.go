package Stack

/*
	LIFO
	Last
	In
	First
	Out
*/

type Stack struct {
	items []int
}

// Push add a value on last index
func (s *Stack) Push(n int) {
	s.items = append(s.items, n) // Appeding on last index
}

// Pop remove a value on last index
func (s *Stack) Pop() int {
	last := len(s.items) - 1 // Getting last index
	removed := s.items[last]
	s.items = s.items[:last] // Slice without last item

	return removed // Returning value removed
}
