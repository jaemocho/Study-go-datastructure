package sort

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	assert := assert.New(t)

	arr := make([]int, 100)
	for i := 0; i < 100; i++ {
		arr[i] = rand.Intn(100)
	}

	assert.False(IsSorted(arr))
	QuickSort(arr)
	assert.True(IsSorted(arr), arr)
	t.Log(arr)
}
