package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPush(t *testing.T) {
	q := New[int]()
	assert := assert.New(t)
	q.Push(1)
	q.Push(2)
	q.Push(3)

	assert.Equal(1, q.Pop())
	assert.Equal(2, q.Pop())
	assert.Equal(3, q.Pop())

}

func TestPush2(t *testing.T) {
	q := NewSliceQueue[int]()
	assert := assert.New(t)
	q.Push(1)
	q.Push(2)
	q.Push(3)

	assert.Equal(1, q.Pop())
	assert.Equal(2, q.Pop())
	assert.Equal(3, q.Pop())

}

func BenchmarkLinkedListQueue(b *testing.B) {
	q := New[int]()
	for i := 0; i < b.N; i++ {
		q.Push(i)
	}
	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}

func BenchmarkLinkedListSliceQueue(b *testing.B) {
	q := NewSliceQueue[int]()
	for i := 0; i < b.N; i++ {
		q.Push(i)
	}
	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}
