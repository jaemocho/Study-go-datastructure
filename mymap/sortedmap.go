package mymap

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// key 로 정렬이 되어야해서 key는 ordered로
type Element[TKey constraints.Ordered, TValue any] struct {
	Key   TKey
	Value TValue
}

type SortedMap[TKey constraints.Ordered, TValue any] struct {
	Arr []Element[TKey, TValue]
}

func (s *SortedMap[TKey, TValue]) Add(key TKey, value TValue) {

	// 삽입할 idx get
	idx := sort.Search(len(s.Arr), func(i int) bool {
		return s.Arr[i].Key >= key
	})

	// sort.Search 가 못찾으면 0을 반환하여 idx < len(s.Arr) 처리 필요
	if idx < len(s.Arr) && s.Arr[idx].Key == key {
		s.Arr[idx].Value = value
		return
	}

	// s.Arr idx 전까지 data로 시작
	s.Arr = append(s.Arr[:idx],
		// 새로운 slice 생성하여 추가할 값 하나 넣고
		append([]Element[TKey, TValue]{{Key: key, Value: value}},
			//s.Arr 의 idx 뒤부터 나머지 값들을 append
			s.Arr[idx:]...)...)

	// sort.Search
	// func Search(n int, f func(int) bool) int {
	// 	// Define f(-1) == false and f(n) == true.
	// 	// Invariant: f(i-1) == false, f(j) == true.
	// 	i, j := 0, n
	// 	for i < j {
	// 		h := int(uint(i+j) >> 1) // avoid overflow when computing h
	// 		// i ≤ h < j
	// 		if !f(h) {    << f 부분을 인자로 받는다 즉 정렬 기준
	// 			i = h + 1 // preserves f(i-1) == false
	// 		} else {
	// 			j = h // preserves f(j) == true
	// 		}
	// 	}
}

func (s *SortedMap[TKey, TValue]) Get(key TKey) (TValue, bool) {
	// 반환할 idx get
	idx := sort.Search(len(s.Arr), func(i int) bool {
		return s.Arr[i].Key >= key
	})

	// return s.Arr[i].Key >= key 방식으로 조회를 해서 같은지 확인 필요
	if idx < len(s.Arr) && s.Arr[idx].Key == key {
		return s.Arr[idx].Value, true
	}

	var defaultV TValue
	return defaultV, false

}

func (s *SortedMap[TKey, TValue]) Remove(key TKey) bool {
	// 삭제할 idx get
	idx := sort.Search(len(s.Arr), func(i int) bool {
		return s.Arr[i].Key >= key
	})

	// return s.Arr[i].Key >= key 방식으로 조회를 해서 같은지 확인 필요
	if idx < len(s.Arr) && s.Arr[idx].Key == key {
		s.Arr = append(s.Arr[:idx], s.Arr[idx+1:]...)
		return true
	}
	return false

}
