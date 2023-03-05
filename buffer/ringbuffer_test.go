package buffer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRingWrite(t *testing.T) {
	assert := assert.New(t)
	buf := NewRingBuffer[byte](10)

	buf.Write([]byte{1, 2, 3, 4})

	assert.Equal(4, buf.Readable())
}

func TestRingRead(t *testing.T) {
	assert := assert.New(t)
	buf := NewRingBuffer[byte](10)

	buf.Write([]byte{1, 2, 3, 4})
	assert.Equal(4, buf.Readable())

	readedData := buf.Read(4)
	for i := 0; i < 4; i++ {
		assert.Equal(byte(i+1), readedData[i])
	}

	assert.Equal(0, buf.Readable())

}

func TestRingBufferOverwrite(t *testing.T) {
	assert := assert.New(t)

	buf := NewRingBuffer[byte](5)
	buf.Write([]byte{1, 2, 3, 4})
	assert.Equal(4, buf.Readable())
	assert.Equal(1, buf.Writable())

	buf.Write([]byte{5})
	assert.Equal(5, buf.Readable())
	assert.Equal(0, buf.Writable())

	writed := buf.Write([]byte{6})
	assert.Equal(0, writed)
	assert.Equal(5, buf.Readable())
	assert.Equal(0, buf.Writable())

	readedData := buf.Read(4)
	for i := 0; i < 4; i++ {
		assert.Equal(byte(i+1), readedData[i])
	}

	assert.Equal(1, buf.Readable())
	assert.Equal(4, buf.Writable())

	writed = buf.Write([]byte{6, 7, 8})
	assert.Equal(3, writed)
	assert.Equal(3, buf.writePoint)
	assert.Equal(4, buf.Readable())
	assert.Equal(1, buf.Writable())

	writed = buf.Write([]byte{6, 7, 8})
	assert.Equal(1, writed)
	assert.Equal(5, buf.Readable())
	assert.Equal(0, buf.Writable())

	readedData = buf.Read(4)
	assert.Equal(4, len(readedData))
	assert.Equal(1, buf.Readable())
	assert.Equal(4, buf.Writable())
}

// func TestWriteAndRest(t *testing.T) {
// 	assert := assert.New(t)
// 	buf := NewSliceBuffer[byte]()

// 	buf.Write([]byte{1, 2, 3, 4})
// 	assert.Equal(4, buf.Readable())

// 	buf.Write([]byte{5, 6})
// 	assert.Equal(6, buf.Readable())

// 	readedData := buf.Read(4)
// 	for i := 0; i < 4; i++ {
// 		assert.Equal(byte(i+1), readedData[i])
// 	}
// 	assert.Equal(2, buf.Readable())

// 	buf.Write([]byte{7, 8, 9})
// 	assert.Equal(5, buf.Readable())

// 	readedData = buf.Read(3)
// 	for i := 0; i < 3; i++ {
// 		assert.Equal(byte(i+5), readedData[i])
// 	}
// 	assert.Equal(2, buf.Readable())

// }
