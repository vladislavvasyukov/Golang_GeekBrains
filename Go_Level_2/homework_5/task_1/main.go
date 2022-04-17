package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const count = 10

func main() {
	var (
		counter int32
		wg      = sync.WaitGroup{}
	)

	wg.Add(count)
	for i := 0; i < count; i += 1 {
		go func(number int) {
			atomic.AddInt32(&counter, 1)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println(counter)
}
