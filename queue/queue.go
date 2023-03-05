package queue

import "goDataStructure/LinkedList/doubleLinkedList"

type Queue[T any] struct {
	l *doubleLinkedList.LinkedList[T]
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		l: &doubleLinkedList.LinkedList[T]{},
	}
}

func (q Queue[T]) Push(val T) {
	q.l.PushBack(val)

}

func (q Queue[T]) Pop() T {
	return q.l.PopFront().Value
}
