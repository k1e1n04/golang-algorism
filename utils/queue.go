package utils

type Queue[T any] struct {
	max  int
	tail int
	head int
	qu   []T
}

func NewQueue[T any]() *Queue[T] {
	qMax := 100000
	return &Queue[T]{
		max:  qMax,
		head: 0,
		tail: 0,
		qu:   make([]T, qMax),
	}
}

// isEmpty はスタックが空かどうか
func (q *Queue[T]) isEmpty() bool {
	return q.head == q.tail
}

// isFull はスタックが満杯かどうか
func (q *Queue[T]) isFull() bool {
	return q.head == (q.tail+1)%q.max
}

// Enqueue は要素を追加
func (q *Queue[T]) Enqueue(v T) {
	if q.isFull() {
		println("queue is full")
		return
	}
	q.qu[q.tail] = v
	q.tail++
	if q.tail == q.max {
		q.tail = 0
	}
}

// Dequeue は要素を取り出し
func (q *Queue[T]) Dequeue() *T {
	if q.isEmpty() {
		println("queue is empty")
		return nil
	}
	res := q.qu[q.head]
	q.head++
	if q.head == q.max {
		q.head = 0
	}
	return &res
}
