package sort

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	assert := assert.New(t)

	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(100)
	}

	assert.False(IsSorted(arr))
	sorted := MergeSort(arr)
	assert.True(IsSorted(sorted), arr)
	t.Log(sorted)
}

func BenchmarkQuickSort(b *testing.B) {
	arr := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		arr[i] = rand.Intn(b.N)
	}

	QuickSort(arr)
}

func BenchmarkMergeSort(b *testing.B) {
	arr := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		arr[i] = rand.Intn(b.N)
	}

	MergeSort(arr)
}

/*
goos: windows
goarch: amd64
pkg: goDataStructure/sort
cpu: Intel(R) Core(TM) i5-7200U CPU @ 2.50GHz
BenchmarkQuickSort-4   	 9032049	       148.0 ns/op	       8 B/op	       0 allocs/op
BenchmarkMergeSort-4   	 6212253	       210.7 ns/op	     190 B/op	       1 allocs/op
PASS
ok  	goDataStructure/sort	3.056s


*/
