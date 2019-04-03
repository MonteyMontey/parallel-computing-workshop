package stack

import (
	"sync"
)

type ThreadSafeStack struct {
	Numbers []int
	maxSize int
	mutex   sync.Mutex
}

func NewThreadSafeStack(maxSize int) *ThreadSafeStack {
	return &ThreadSafeStack{
		maxSize: maxSize,
	}
}

func (s *ThreadSafeStack) Push(number int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.Numbers) < s.maxSize {
		s.Numbers = append(s.Numbers, number) // only locking this line would lead to race condition
	}
}

func (s *ThreadSafeStack) Pop() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.Numbers) > 0 {
		s.Numbers = s.Numbers[:len(s.Numbers)-1] // only locking this line would lead to race condition
	}
}

func (s *ThreadSafeStack) Empty() {
	s.mutex.Lock()
	defer s.mutex.Unlock()


	if len(s.Numbers) > 0 {
		s.Numbers = nil // only locking this line would lead to race condition
	}
}
