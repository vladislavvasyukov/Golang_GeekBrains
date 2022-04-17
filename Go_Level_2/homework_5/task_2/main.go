package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		mu      sync.Mutex
		counter = 10
		wg      = sync.WaitGroup{}
	)

	wg.Add(counter)

	for i := 0; i < counter; i++ {
		go func(number int) {
			mu.Lock()
			defer mu.Unlock()
			fmt.Println(number)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("The end...")
}
