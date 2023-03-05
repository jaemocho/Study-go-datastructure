package buffer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWrite(t *testing.T) {
	assert := assert.New(t)
	buf := NewSliceBuffer[byte]()

	buf.Write([]byte{1, 2, 3, 4})

	assert.Equal(4, buf.Readable())
}

func TestRead(t *testing.T) {
	assert := assert.New(t)
	buf := NewSliceBuffer[byte]()

	buf.Write([]byte{1, 2, 3, 4})
	assert.Equal(4, buf.Readable())

	readedData := buf.Read(4)
	for i := 0; i < 4; i++ {
		assert.Equal(byte(i+1), readedData[i])
	}

	assert.Equal(0, buf.Readable())

}

func TestWriteAndRest(t *testing.T) {
	assert := assert.New(t)
	buf := NewSliceBuffer[byte]()

	buf.Write([]byte{1, 2, 3, 4})
	assert.Equal(4, buf.Readable())

	buf.Write([]byte{5, 6})
	assert.Equal(6, buf.Readable())

	readedData := buf.Read(4)
	for i := 0; i < 4; i++ {
		assert.Equal(byte(i+1), readedData[i])
	}
	assert.Equal(2, buf.Readable())

	buf.Write([]byte{7, 8, 9})
	assert.Equal(5, buf.Readable())

	readedData = buf.Read(3)
	for i := 0; i < 3; i++ {
		assert.Equal(byte(i+5), readedData[i])
	}
	assert.Equal(2, buf.Readable())

}
