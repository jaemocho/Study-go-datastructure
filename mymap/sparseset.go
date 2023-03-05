package mymap

import "golang.org/x/exp/constraints"

// SparseSet 은 순서보장은 없다. 저장만 dense 하게

/*
삽입/삭제/조회를 빠르게
캐쉬 친화적
Dense와 sparse 자료구조를 둘다 사용
key를 slice에 저장하는 방식

map 에 데이터의 idx 위치를 저장 하여 사용
데이터는 arr 에 저장 (빈공간 없이)

iterator가 빨라야할 때 사용
게임 업계에서 많은 entity를 사용하는 상황에서 사용
*/

type SparseSet[TKey constraints.Ordered, TValue any] struct {
	// array 에 데이터를 보관함으로서 dense 하게 관리
	dense  []Element[TKey, TValue]
	sparse map[TKey]int
}

func NewSparseSet[TKey constraints.Ordered, TValue any]() *SparseSet[TKey, TValue] {
	return &SparseSet[TKey, TValue]{
		// map이 reference type이라 선언하지 않으면 nil porinter access 발생
		sparse: make(map[TKey]int),
	}
}

func (s *SparseSet[TKey, TValue]) Add(key TKey, value TValue) {
	if idx, ok := s.sparse[key]; ok {
		s.dense[idx].Value = value
		return
	}

	s.dense = append(s.dense, Element[TKey, TValue]{
		Key:   key,
		Value: value,
	})

	s.sparse[key] = len(s.dense) - 1
}

func (s *SparseSet[TKey, TValue]) Get(key TKey) (value TValue, found bool) {
	if idx, ok := s.sparse[key]; ok {
		value = s.dense[idx].Value
		found = true
		return
	}

	found = false
	return
}

//1. dense 에 마지막 idx의 값을 삭제되는 곳으료 이동 (빈공간이 없게)
//2. map에서 idx 삭제
func (s *SparseSet[TKey, TValue]) Remove(key TKey) bool {
	if idx, ok := s.sparse[key]; ok {
		last := len(s.dense) - 1
		// idx 가 last이면 삭제만 진행해주면 되어서 이동은 필요 없음
		if idx < last {
			// 배열의 마지막 값을 삭제되는 곳으로 복사
			s.dense[idx] = s.dense[last]
			// 위에서 복사 했으니 map에서 key 위치(배열의 idx)를 변경
			s.sparse[s.dense[last].Key] = idx
		}
		s.dense = s.dense[:last]
		delete(s.sparse, key)
		return true

	}

	return false
}

type Iterator[TKey constraints.Ordered, TValue any] struct {
	dense []Element[TKey, TValue]
	idx   int
}

func (s *SparseSet[TKey, TValue]) Iterator() *Iterator[TKey, TValue] {
	return &Iterator[TKey, TValue]{
		// slice는 len cap, startpoint 만 copy 된다
		dense: s.dense,
		idx:   0,
	}
}

func (i *Iterator[TKey, TValue]) IsEnd() bool {
	return i.idx >= len(i.dense)
}

func (i *Iterator[TKey, TValue]) Next() {
	i.idx++
}

func (i *Iterator[TKey, TValue]) Get() Element[TKey, TValue] {
	return i.dense[i.idx]
}
