package LinkedList

import (
	"goDataStructure/LinkedList/doubleLinkedList"
	"goDataStructure/LinkedList/singleLinkedList"
	"testing"
)

func BenchmarkSingleLinkedList(b *testing.B) {
	var l singleLinkedList.LinkedList[int]

	for i := 0; i < b.N; i++ {
		l.PushBack(i)
	}

	l.Reverse()
}

func BenchmarkSingleLinkedListSwap(b *testing.B) {
	var l singleLinkedList.LinkedList[int]

	for i := 0; i < b.N; i++ {
		l.PushBack(i)
	}

	l.Reverse2()
}

func BenchmarkDoubleLinkedList(b *testing.B) {
	var l doubleLinkedList.LinkedList[int]
	for i := 0; i < b.N; i++ {
		l.PushBack(i)
	}

	l.Reverse()
}

/*

goos: windows
goarch: amd64
pkg: linkedList
cpu: Intel(R) Core(TM) i5-7200U CPU @ 2.50GHz
BenchmarkSingleLinkedList-4       	 9952575	       133.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkSingleLinkedListSwap-4   	19548080	        90.05 ns/op	      16 B/op	       1 allocs/op
BenchmarkDoubleLinkedList-4       	11383966	       103.6 ns/op	      24 B/op	       1 allocs/op
PASS
ok  	linkedList	4.711s


*/
