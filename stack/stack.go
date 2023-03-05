package stack

import "goDataStructure/LinkedList/doubleLinkedList"

type Stack[T any] struct {
	l *doubleLinkedList.LinkedList[T]
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		l: &doubleLinkedList.LinkedList[T]{},
	}
}

func (s Stack[T]) Push(val T) {
	s.l.PushBack(val)

}

func (s Stack[T]) Pop() T {
	return s.l.PopBack().Value
}
