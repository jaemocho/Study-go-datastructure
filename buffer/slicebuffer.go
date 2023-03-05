package buffer

// array 의 길이가 계속 바뀐다.
// 새로운 array가 allocation 되는 현상이 계속 발생
// gc memory 사용량이 늘어난다. 성능에 문제가 발생
// Ring buffer 사용하는 것이 더 좋다

type SliceBuffer[T any] struct {
	buf []T
}

func NewSliceBuffer[T any]() *SliceBuffer[T] {
	return &SliceBuffer[T]{}
}

func (s *SliceBuffer[T]) Write(data []T) {
	// data 가 slice type이라 data... 으로 append
	s.buf = append(s.buf, data...)
}

func (s *SliceBuffer[T]) Readable() int {
	return len(s.buf)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func (s *SliceBuffer[T]) Read(count int) []T {
	// 요청한 개수와 현재 버퍼의 남은 개수 중 작은 수만큼만 읽는다.
	readCnt := min(count, s.Readable())

	rst := make([]T, readCnt)

	copy(rst, s.buf[:readCnt])
	// 읽어진 부분 제거
	s.buf = s.buf[readCnt:]
	return rst
}
