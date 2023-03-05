package sort

import "golang.org/x/exp/constraints"

// constraints.Ordered 대소 비교가 가능한 타입들만
func QuickSort[T constraints.Ordered](arr []T) {
	if len(arr) <= 1 {
		return
	}
	//idx 다음 pivot 의 위치
	idx := partition(arr)
	// pivot 기준 앞부분 정렬
	QuickSort(arr[:idx])
	// pivot 기준 뒷부분 정렬
	QuickSort(arr[idx+1:])
}

func partition[T constraints.Ordered](arr []T) int {
	if len(arr) <= 1 {
		return 0
	}
	// 처음 값 마지막 값 중간 값 중 골라서 사용하는데
	// 여기선 처음 값으로 구현
	pivot := arr[0]
	// pivot 기준으로 작으면 왼쪽 크면 오른쪽
	i := 1
	j := len(arr) - 1
	for {
		// arr[i] 의 값이 arr[0] 즉 pivot의 값보다 작은 것을 찾는다.
		for i < len(arr) && arr[i] <= pivot {
			i++
		}
		// arr[j] 의 값이 arr[0] 즉 pivot의 값보다 큰 것을 찾는다.
		for j > 0 && arr[j] > pivot {
			j--
		}

		// 두 idx가 역전되는 순간이 다음 pivot의 위치
		if i >= j {
			// 현재 pivot 값과 처음 정한 pivot의 값 swap
			arr[0], arr[i-1] = arr[i-1], arr[0]
			return i - 1
		}
		// 위의 경우가 아닌경우 두 값만 바꿔주고 다시 진행
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func IsSorted[T constraints.Ordered](arr []T) bool {
	if len(arr) <= 1 {
		return true
	}

	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}

	return true
}
