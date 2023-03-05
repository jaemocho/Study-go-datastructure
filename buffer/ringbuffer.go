package buffer

type RingBuffer[T any] struct {
	buf        []T
	readPoint  int
	writePoint int
	// readPoint, writePoint가 같은 곳에 있는 경우 데이터가 다 차있는지 비어있는 지 확인할 방법이 없어서 flag 사용
	isFull bool
}

func NewRingBuffer[T any](size int) *RingBuffer[T] {
	return &RingBuffer[T]{
		buf: make([]T, size),
	}
}

func (r *RingBuffer[T]) Readable() int {
	if r.isFull {
		return len(r.buf)
	}
	if r.writePoint >= r.readPoint {
		return r.writePoint - r.readPoint
	} else {
		return len(r.buf) - r.readPoint + r.writePoint
	}
}

func (r *RingBuffer[T]) Writable() int {
	return len(r.buf) - r.Readable()
}

func (r *RingBuffer[T]) Write(data []T) int {
	if len(data) == 0 || r.Writable() == 0 {
		return 0
	}
	var writed int
	if r.writePoint > r.readPoint { // read 와 write 사이에 채워져 있는 상황

		//1. writePoint 부터 end 까지 쓸 수 있는만큼 먼저 쓰는 구간
		writableToEnd := len(r.buf) - r.writePoint
		writed = min(writableToEnd, len(data))

		copy(r.buf[r.writePoint:], data[:writed])
		r.writePoint += writableToEnd

		// 앞에서부터 readPoint 까지의 거리를 계산하기 위한  나머지 연산
		r.writePoint %= len(r.buf)

		// 2. 0 부터 readPoint 까지 쓸 수 있는 만큼 쓰는 구간
		// 재귀로 호출
		remain := len(data) - writed
		if remain > 0 {
			writed += r.Write(data[writed:])
		}

	} else { // 0 ~ read , write ~len 까지 채워져 있는 상황
		writed = min(r.Writable(), len(data))
		copy(r.buf[r.writePoint:], data[:writed])
		r.writePoint += writed
		// 앞에서부터 readPoint 까지의 거리를 계산하기 위한  나머지 연산
		r.writePoint %= len(r.buf)

	}

	if writed > 0 && r.writePoint == r.readPoint {
		r.isFull = true
	}
	return writed

}

func (r *RingBuffer[T]) Read(count int) []T {
	if r.Readable() == 0 || count <= 0 {
		return []T{}
	}
	readCnt := min(count, r.Readable())
	rst := make([]T, readCnt)

	// 배열 끝을 넘어가는 경우 처리
	if r.readPoint+readCnt >= len(r.buf) {
		remainUntilEnd := len(r.buf) - r.readPoint
		// 1. 배열 끝까지 읽어준다.
		copy(rst, r.buf[r.readPoint:])
		r.readPoint = 0

		// 2. 0에서 부터 남은 개수 만큼 읽어준다.
		remain := readCnt - remainUntilEnd
		copy(rst[remainUntilEnd:], r.buf[:remain])
		r.readPoint += remain
	} else {
		copy(rst, r.buf[r.readPoint:r.readPoint+readCnt])
		r.readPoint += readCnt
	}

	// 읽었다는 거니까 false 처리
	r.isFull = false
	return rst
}
