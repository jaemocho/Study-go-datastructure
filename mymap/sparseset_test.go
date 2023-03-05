package mymap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSparseSet(t *testing.T) {
	assert := assert.New(t)
	s := NewSparseSet[string, int]()

	s.Add("ccc", 30)
	s.Add("bbb", 20)
	s.Add("aaa", 10)

	v, ok := s.Get("bbb")
	assert.True(ok)
	assert.Equal(20, v)

	v, ok = s.Get("aaa")
	assert.True(ok)
	assert.Equal(10, v)

	v, ok = s.Get("ccc")
	assert.True(ok)
	assert.Equal(30, v)

	_, ok = s.Get("ddd")
	assert.False(ok)

	removed := s.Remove("bbb")
	assert.True(removed)

	_, ok = s.Get("bbb")
	assert.False(ok)

	for i := s.Iterator(); !i.IsEnd(); i.Next() {
		element := i.Get()
		if element.Key == "aaa" {
			assert.Equal(10, element.Value)
		} else if element.Key == "ccc" {
			assert.Equal(30, element.Value)
		} else {
			t.Fail()
		}

	}
}
