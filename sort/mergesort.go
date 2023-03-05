package sort

import (
	"golang.org/x/exp/constraints"
)

func MergeSort[T constraints.Ordered](arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)

}

func merge[T constraints.Ordered](left, right []T) []T {
	i := 0
	j := 0
	idx := 0
	// merge 할 슬라이스 생성
	rst := make([]T, len(left)+len(right))

	// i 는 왼쪽 배열 j 는 오른쪽 배열의 idx
	for i < len(left) || j < len(right) {
		var leftMerge bool
		// i가 범위가 넘어섰을 경우 오른쪽 것만 merge
		if i >= len(left) {
			leftMerge = false
			// j가 범위가 넘어섰을 경우 왼쪽 것만 merge
		} else if j >= len(right) {
			leftMerge = true
			// i,j둘다 범위 내에 있을 땐 값을 비교
			// left[i] 가 작으면 true leftMerge
			// right[j] 가 작거나 같으면 false rightMerge
		} else {
			leftMerge = left[i] < right[j]
		}

		// merge할 slice에 값을 넣고 idx 증가
		if leftMerge {
			rst[idx] = left[i]
			i++
		} else {
			rst[idx] = right[j]
			j++
		}
		idx++
	}
	return rst
}
