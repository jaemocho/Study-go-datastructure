package mymap

import "hash/crc32"

const arraySize = 3571

type HashData[T any] struct {
	key   string
	value T
}

// spase한 자료구조 듬성듬성, 캐시에 비친화적
type HashMap[T any] struct {
	// hashdata 충돌방지를 위해 slice 형태로 HashData 를 선언
	arr [arraySize][]HashData[T]
}

// hashing 하기 위한 fn
// crc-32 checksum 용 단순한 알고리즘
func hashfn(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

func (h *HashMap[T]) Add(key string, value T) {
	hash := hashfn(key)
	var hd = HashData[T]{
		key:   key,
		value: value,
	}
	h.arr[hash%arraySize] = append(h.arr[hash%arraySize], hd)
}

func (h *HashMap[T]) Get(key string) (T, bool) {
	hash := hashfn(key)
	for _, hd := range h.arr[hash%arraySize] {
		if hd.key == key {
			return hd.value, true
		}
	}
	var tempT T
	return tempT, false
}
