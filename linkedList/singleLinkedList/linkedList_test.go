package singleLinkedList

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushBack(t *testing.T) {
	var l LinkedList[int]

	assert := assert.New(t)

	assert.Nil(l.root)
	assert.Nil(l.tail)
	l.PushBack(1)

	assert.NotNil(l.root)
	assert.Equal(1, l.Front().Value)
	assert.Equal(1, l.Back().Value)

	l.PushBack(2)

	assert.NotNil(l.root)
	assert.Equal(1, l.Front().Value)
	assert.Equal(2, l.Back().Value)

	l.PushBack(3)

	assert.NotNil(l.root)
	assert.Equal(1, l.Front().Value)
	assert.Equal(3, l.Back().Value)

	assert.Equal(3, l.Count())
	assert.Equal(3, l.Count2())

	assert.Equal(1, l.GetAt(0).Value)
	assert.Equal(2, l.GetAt(1).Value)
	assert.Equal(3, l.GetAt(2).Value)
	assert.Nil(l.GetAt(3))
}

func TestPushFront(t *testing.T) {
	var l LinkedList[int]

	assert := assert.New(t)

	assert.Nil(l.root)
	assert.Nil(l.tail)
	l.PushFront(1)

	assert.NotNil(l.root)
	assert.Equal(1, l.Front().Value)
	assert.Equal(1, l.Back().Value)

	l.PushFront(2)

	assert.NotNil(l.root)
	assert.Equal(2, l.Front().Value)
	assert.Equal(1, l.Back().Value)

	l.PushFront(3)

	assert.NotNil(l.root)
	assert.Equal(3, l.Front().Value)
	assert.Equal(1, l.Back().Value)

	assert.Equal(3, l.Count())
	assert.Equal(3, l.Count2())

	l.PushFront(4)

	assert.NotNil(l.root)
	assert.Equal(4, l.Front().Value)
	assert.Equal(1, l.Back().Value)

	assert.Equal(4, l.Count())
	assert.Equal(4, l.Count2())
}

func TestInsertAfter(t *testing.T) {
	assert := assert.New(t)

	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	node := l.GetAt(1)
	l.InsertAfter(node, 4)

	assert.Equal(4, l.Count2())
	assert.Equal(4, l.GetAt(2).Value)
	assert.Equal(3, l.Back().Value)

	tempNode := &Node[int]{
		Value: 100,
	}

	l.InsertAfter(tempNode, 200)
	assert.Equal(4, l.Count2())
}

func TestInsertBefore(t *testing.T) {
	assert := assert.New(t)

	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	node := l.GetAt(1)
	l.InsertBefore(node, 4)

	assert.Equal(4, l.Count2())
	assert.Equal(4, l.GetAt(1).Value)
	assert.Equal(2, l.GetAt(2).Value)
	assert.Equal(3, l.Back().Value)

	tempNode := &Node[int]{
		Value: 100,
	}

	l.InsertBefore(tempNode, 200)
	assert.Equal(4, l.Count2())

}

func TestInsertBeforeRoot(t *testing.T) {
	assert := assert.New(t)

	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	l.InsertBefore(l.GetAt(0), 4)

	assert.Equal(4, l.Count2())
	assert.Equal(1, l.GetAt(1).Value)
	assert.Equal(2, l.GetAt(2).Value)
	assert.Equal(3, l.Back().Value)

	tempNode := &Node[int]{
		Value: 100,
	}

	l.InsertBefore(tempNode, 200)
	assert.Equal(4, l.Count2())

}

func TestPopFront(t *testing.T) {
	assert := assert.New(t)

	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	l.PopFront()
	assert.Equal(2, l.Count2())
	assert.Equal(2, l.Front().Value)
	assert.Equal(3, l.Back().Value)

	l.PopFront()
	assert.Equal(1, l.Count2())
	assert.Equal(3, l.Front().Value)
	assert.Equal(3, l.Back().Value)

	l.PopFront()
	assert.Equal(0, l.Count2())
	assert.Nil(l.Front())
	assert.Nil(l.Back())

}

func TestRemove(t *testing.T) {
	assert := assert.New(t)

	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	l.Remove(l.GetAt(1))
	assert.Equal(2, l.Count2())
	assert.Equal(1, l.Front().Value)
	assert.Equal(3, l.Back().Value)

	// 첫번째 node 삭제 테스트
	l.Remove(l.GetAt(0))
	assert.Equal(1, l.Count2())
	assert.Equal(3, l.Front().Value)
	assert.Equal(3, l.Back().Value)

	// 없는 node 삭제 테스트
	l.Remove(&Node[int]{
		Value: 100,
	})

	assert.Equal(1, l.Count2())
	assert.Equal(3, l.Front().Value)
	assert.Equal(3, l.Back().Value)

	// 하나 남은 node 삭제 테스트
	l.Remove(l.GetAt(0))
	assert.Equal(0, l.Count2())
	assert.Nil(l.Front())
	assert.Nil(l.Back())

	// 마지막 삭제 테스트
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.Remove(l.GetAt(2))
	assert.Equal(2, l.Back().Value)

}

func TestReverse(t *testing.T) {

	assert := assert.New(t)

	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.Reverse()

	assert.Equal(4, l.Front().Value)
	assert.Equal(1, l.Back().Value)
	assert.Equal(3, l.GetAt(1).Value)

}

func TestReverse2(t *testing.T) {

	assert := assert.New(t)

	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.Reverse2()

	assert.Equal(4, l.Front().Value)
	assert.Equal(1, l.Back().Value)
	assert.Equal(3, l.GetAt(1).Value)

	l.PushBack(5)
	l.Reverse2()

	assert.Equal(5, l.Front().Value)
	assert.Equal(4, l.Back().Value)
	assert.Equal(1, l.GetAt(1).Value)
}
