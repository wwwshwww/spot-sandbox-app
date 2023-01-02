package common

import "container/list"

type Deque[T any] struct {
	li *list.List
}

func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{li: list.New()}
}

func (q *Deque[T]) Append(n T) {
	q.li.PushBack(n)
}

func (q *Deque[T]) AppendLeft(n T) {
	q.li.PushFront(n)
}

func (q *Deque[T]) Pop() T {
	e := q.li.Back()
	defer func() {
		q.li.Remove(e)
	}()
	return e.Value.(T)
}

func (q *Deque[T]) PopLeft() T {
	e := q.li.Front()
	defer func() {
		q.li.Remove(e)
	}()
	return e.Value.(T)
}

func (q *Deque[T]) Len() int {
	return q.li.Len()
}
