package singleLinkedList

type Node[T any] struct {
	Next  *Node[T]
	Value T
}

type LinkedList[T any] struct {
	root  *Node[T]
	tail  *Node[T]
	count int
}

func (l *LinkedList[T]) PushBack(value T) {
	node := &Node[T]{
		Value: value,
	}
	l.count++

	if l.root == nil {
		l.root = node
		l.tail = node
		return
	}
	l.tail.Next = node
	l.tail = node
}

func (l *LinkedList[T]) PushFront(value T) {
	node := &Node[T]{
		Value: value,
	}
	l.count++
	if l.root == nil {
		l.root = node
		l.tail = node
		return
	}
	node.Next = l.root
	l.root = node
}

func (l *LinkedList[T]) Front() *Node[T] {
	return l.root
}

func (l *LinkedList[T]) Back() *Node[T] {
	return l.tail
}

func (l *LinkedList[T]) Count() int {
	node := l.root
	cnt := 0

	for ; node != nil; node = node.Next {
		cnt++
	}

	return cnt
}

func (l *LinkedList[T]) Count2() int {
	return l.count
}

func (l *LinkedList[T]) GetAt(idx int) *Node[T] {
	if idx >= l.Count2() {
		return nil
	}

	i := 0
	for node := l.root; node != nil; node = node.Next {
		if i == idx {
			return node
		}
		i++
	}

	return nil
}

func (l *LinkedList[T]) InsertAfter(node *Node[T], value T) {
	if !l.isIncluded(node) {
		return
	}
	newNode := &Node[T]{
		Value: value,
	}

	originNext := node.Next
	node.Next = newNode
	newNode.Next = originNext

	//node.Next, newNode.Next = newNode, node.Next
	l.count++
}

func (l *LinkedList[T]) isIncluded(node *Node[T]) bool {
	inner := l.root
	for ; inner != nil; inner = inner.Next {
		if inner == node {
			return true
		}
	}

	return false
}

func (l *LinkedList[T]) InsertBefore(node *Node[T], value T) {
	if node == l.root {
		l.PushFront(value)
		return
	}
	prevNode := l.findPrevNode(node)
	if prevNode == nil {
		return
	}
	newNode := &Node[T]{
		Value: value,
	}

	newNode.Next = node
	prevNode.Next = newNode
	l.count++
}

func (l *LinkedList[T]) findPrevNode(node *Node[T]) *Node[T] {
	inner := l.root
	for ; inner != nil; inner = inner.Next {
		if inner.Next == node {
			return inner
		}
	}

	return nil
}

func (l *LinkedList[T]) PopFront() *Node[T] {
	if l.root == nil {
		return nil
	}
	n := l.root
	temp := l.root.Next
	l.root.Next = nil
	l.root = temp

	if l.root == nil {
		l.tail = nil
	}

	l.count--
	return n
}

func (l *LinkedList[T]) Remove(node *Node[T]) {
	if node == l.root {
		l.PopFront()
		return
	}

	prev := l.findPrevNode(node)
	if prev == nil {
		return
	}
	prev.Next = node.Next
	node.Next = nil

	// ????????? node ?????? ??? ??????
	if node == l.tail {
		l.tail = prev
	}

	l.count--
}

// pop & push ????????? ????????? ?????? ????????? ??????
func (l *LinkedList[T]) Reverse() {
	if l.root == nil {
		return
	}
	newL := &LinkedList[T]{}

	for l.root != nil {
		n := l.PopFront()
		newL.PushFront(n.Value)
	}

	l.count = newL.count
	l.root = newL.root
	l.tail = newL.tail
}

// 2?????? ???????????? ????????? root ??? ??????
func (l *LinkedList[T]) Reverse2() {
	if l.root == nil {
		return
	}

	// ????????? link ???????????? ???????????? ??????
	cnt := l.count

	// ????????? ???????????? ??????
	var temp *Node[T] = l.root.Next
	var next *Node[T]

	l.tail = l.root
	for i := 1; i < cnt; i++ {
		// ????????? ?????? ??????????????? ?????? ??????
		next = temp.Next

		// ?????? ?????? ????????? ?????? ?????????
		temp.Next = l.root
		// root ??????
		l.root = temp

		// ?????? ?????????????????? node??? ?????? ??????
		temp = next
	}
	l.tail.Next = nil
}
