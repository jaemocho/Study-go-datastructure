package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPush(t *testing.T) {
	s := New[int]()
	assert := assert.New(t)
	s.Push(1)
	s.Push(2)
	s.Push(3)

	assert.Equal(3, s.Pop())
	assert.Equal(2, s.Pop())
	assert.Equal(1, s.Pop())

}

func TestPush2(t *testing.T) {
	s := NewSliceStack[int]()
	assert := assert.New(t)
	s.Push(1)
	s.Push(2)
	s.Push(3)

	assert.Equal(3, s.Pop())
	assert.Equal(2, s.Pop())
	assert.Equal(1, s.Pop())

}

func BenchmarkLinkedListStack(b *testing.B) {
	s := New[int]()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}

func BenchmarkLinkedListSliceStack(b *testing.B) {
	s := NewSliceStack[int]()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}

/*

goos: windows
goarch: amd64
pkg: goDataStructure/Stack
cpu: Intel(R) Core(TM) i5-7200U CPU @ 2.50GHz
BenchmarkLinkedListStack-4        	13149934	        77.76 ns/op	      24 B/op	       1 allocs/op
BenchmarkLinkedListSliceStack-4   	59517609	        25.91 ns/op	      49 B/op	       0 allocs/op
PASS
ok  	goDataStructure/Stack	2.978s

*/
