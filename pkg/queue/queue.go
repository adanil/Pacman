package queue

type Queue[T any] struct {
	arr []T
}

func (q *Queue[T]) Push(el T) {
	q.arr = append(q.arr, el)
}

func (q *Queue[T]) Front() T {
	return q.arr[0]
}

func (q *Queue[T]) Pop() {
	q.arr = q.arr[1:len(q.arr)]
}

func (q *Queue[T]) Size() int {
	return len(q.arr)
}
