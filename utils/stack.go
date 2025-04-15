package utils

// Stack はスタック
type Stack[T any] struct {
	max int
	top int
	st  []T
}

// NewStack はスタックを生成
func NewStack[T any]() *Stack[T] {
	sMax := 1000000
	return &Stack[T]{
		max: sMax,
		top: 0,
		st:  make([]T, sMax),
	}
}

// isEmpty はスタックが空かどうか
func (s *Stack[T]) isEmpty() bool {
	return s.top == 0
}

// isFull はスタックが満杯かどうか
func (s *Stack[T]) isFull() bool {
	return s.top == s.max
}

// Push は追加
func (s *Stack[T]) Push(v T) {
	if s.isFull() {
		println("stack is full")
	}
	s.st[s.top] = v
	s.top++
}

// Pop は最後の要素を取り出し
func (s *Stack[T]) Pop() T {
	if s.isEmpty() {
		println("stack is full")
	}
	s.top--
	return s.st[s.top]
}
