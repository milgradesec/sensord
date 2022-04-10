package serial

import (
	"fmt"
	"math/rand"
	"time"
)

type MockReader struct {
	ch chan string
}

func NewMockReader(ch chan string) *MockReader {
	return &MockReader{
		ch: ch,
	}
}

func (r *MockReader) Start() {
	for {
		time.Sleep(300 * time.Millisecond)

		n := rand.Intn(1023-500) + 500 //nolint
		r.ch <- fmt.Sprintf("%d", n)
	}
}

func (r *MockReader) Close() error {
	return nil
}
