package mymap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedMap(t *testing.T) {
	assert := assert.New(t)
	var s SortedMap[string, int]

	s.Add("aaa", 10)
	v, ok := s.Get("aaa")
	assert.Equal(10, v)
	assert.True(ok)

	s.Add("bbb", 20)
	v, ok = s.Get("bbb")
	assert.Equal(20, v)
	assert.True(ok)

	assert.Equal("aaa", s.Arr[0].Key)
	assert.Equal("bbb", s.Arr[1].Key)
}

func TestSortedMapOverlapped(t *testing.T) {
	assert := assert.New(t)
	var s SortedMap[string, int]

	// 덮어 씌워 주는 test
	s.Add("bbb", 10)
	v, ok := s.Get("bbb")
	assert.Equal(10, v)
	assert.True(ok)

	s.Add("bbb", 20)
	v, ok = s.Get("bbb")
	assert.Equal(20, v)
	assert.True(ok)

	assert.Equal(1, len(s.Arr))

}

func TestSortedGetEmpty(t *testing.T) {
	assert := assert.New(t)
	var s SortedMap[string, int]

	_, ok := s.Get("aaa")
	assert.False(ok)

	s.Add("aaa", 10)
	v, ok := s.Get("aaa")
	assert.Equal(10, v)
	assert.True(ok)

}

func TestSortedMapRemove(t *testing.T) {
	assert := assert.New(t)
	var s SortedMap[string, int]

	s.Add("ccc", 30)
	s.Add("bbb", 20)
	s.Add("aaa", 10)

	removed := s.Remove("bbb")
	assert.True(removed)

	removed = s.Remove("bbb")
	assert.False(removed)

	assert.Equal(2, len(s.Arr))

	assert.Equal("aaa", s.Arr[0].Key)
	assert.Equal("ccc", s.Arr[1].Key)

}
