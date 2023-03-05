package main

import "fmt"

func main() {
	arr := []int{5, 3, 5, 4, 8, 3, 1, 2, 9, 7, 4, 2, 3, 9, 7, 1}

	var count [11]int

	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}

	fmt.Println(len(arr))

	// 1. count 한 후 sorted
	sorted := make([]int, 0, len(arr))
	for i := 0; i < 11; i++ {
		for j := 0; j < count[i]; j++ {
			sorted = append(sorted, i)
		}
	}
	fmt.Println(sorted)

	// 2. count 한 값을 누적 후 출력
	// 중복 for문 사용 안함 1. 번보다 성능이 조금 더 좋다
	for i := 1; i < 11; i++ {
		count[i] += count[i-1]
	}

	sorted = make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		sorted[count[arr[i]]-1] = arr[i]
		count[arr[i]]--
	}
	fmt.Println(count)
	fmt.Println(sorted)
}
