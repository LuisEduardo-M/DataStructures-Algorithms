package Queue

/*
	FIFO
	First
	In
	First
	Out
*/

type Queue struct {
	items []int
}

// Enqueue add the value on last index
func (q *Queue) Enqueue(n int) {
	q.items = append(q.items, n) // Appending on last index
}

// Dequeue remove the value on first index
func (q *Queue) Dequeue() int {
	removed := q.items[0] // Removing first item
	q.items = q.items[1:] // Slice without first item

	return removed // Returning removed item
}
