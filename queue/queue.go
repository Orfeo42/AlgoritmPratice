package queue

type Queue[T any] struct {
	elements []T
}

func (q *Queue[T]) Pop() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}
	val := q.elements[0]
	q.elements = q.elements[1:]
	return val, true
}

func (q *Queue[T]) Peek() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}
	return q.elements[0], true
}

func (q *Queue[T]) Push(newElement T) {
	q.elements = append(q.elements, newElement)
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q *Queue[T]) Size() int {
	return len(q.elements)
}
