package stack

import (
	"sync"
	"testing"
)

func TestThreadSafeStack(t *testing.T){
	var wg sync.WaitGroup

	threadSafeStack := NewThreadSafeStack(1000)

	for i := 1; i < 100; i++{
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			threadSafeStack.Push(i)
		}(i)
	}

	wg.Wait()
}
