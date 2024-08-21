package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type State struct {
	mu      sync.Mutex
	count   int
	count32 int32
}

func (s *State) setCount(i int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.count += i
}

func (s *State) setCountAtomic(i int) {
	atomic.AddInt32(&s.count32, int32(i))
}

func main() {
	state := &State{}
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			state.setCount(i + 1)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("%+v\n", state)
}
