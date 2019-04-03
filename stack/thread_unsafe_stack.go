package stack

type ThreadUnsafeStack struct {
	Numbers []int
	maxSize int
}

func NewThreadUnsafeStack(maxSize int) *ThreadUnsafeStack {
	return &ThreadUnsafeStack{
		maxSize: maxSize,
	}
}

func (s *ThreadUnsafeStack) Push(number int) {
	if len(s.Numbers) < s.maxSize {
		s.Numbers = append(s.Numbers, number)
	}
}

func (s *ThreadUnsafeStack) Pop() {
	if len(s.Numbers) > 0 {
		s.Numbers = s.Numbers[:len(s.Numbers)-1]
	}
}

func (s *ThreadUnsafeStack) Empty() {
	if len(s.Numbers) > 0 {
		s.Numbers = nil
	}
}