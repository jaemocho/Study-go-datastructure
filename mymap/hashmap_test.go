package mymap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashMap(t *testing.T) {
	assert := assert.New(t)
	var h HashMap[int]

	h.Add("jjm", 100)
	val, ok := h.Get("jjm")
	assert.True(ok)
	assert.Equal(100, val)

	h.Add("golang", 200)
	val, ok = h.Get("golang")
	assert.True(ok)
	assert.Equal(200, val)

	h.Add("golangawesome", 300)
	val, ok = h.Get("golangawesome")
	assert.True(ok)
	assert.Equal(300, val)
}

func TestGoBasicMap(t *testing.T) {
	assert := assert.New(t)
	m := make(map[string]int)
	m["jjm"] = 100
	m["golang"] = 200
	m["golangawesome"] = 300

	assert.Equal(100, m["jjm"])
	assert.Equal(200, m["golang"])
	assert.Equal(300, m["golangawesome"])

	// 없는 값 호출 시 default 값 int type이라 0 반환
	// ok(bool) 값으로  판단
	assert.Equal(0, m["aaa"])
	_, ok := m["aaa"]
	assert.False(ok)

	delete(m, "jjm")
	_, ok = m["jjm"]
	assert.False(ok)
}
