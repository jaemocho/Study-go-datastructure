package sort

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertSort(t *testing.T) {
	assert := assert.New(t)
	arr := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		BinaryInsertSort(arr, rand.Intn(100))
	}

	assert.True(IsSorted(arr), arr)
	t.Log(arr)
}

func BenchmarkInsertSort(b *testing.B) {
	arr := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		arr = BinaryInsertSort(arr, rand.Intn(b.N))
	}

}

/*
goos: windows
goarch: amd64
pkg: goDataStructure/sort
cpu: Intel(R) Core(TM) i5-7200U CPU @ 2.50GHz
BenchmarkInsertSort-4   	  147300	    109076 ns/op	  298217 B/op	       1 allocs/op
PASS
ok  	goDataStructure/sort	16.186s

*/
