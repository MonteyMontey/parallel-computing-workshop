package stack

import (
	"sync"
	"testing"
)

func TestThreadUnsafeStack(t *testing.T){
	var wg sync.WaitGroup

	threadUnsafeStack := NewThreadUnsafeStack(1000)

	for i := 1; i < 100; i++{
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			threadUnsafeStack.Push(i)
		}(i)
	}

	wg.Wait()
}