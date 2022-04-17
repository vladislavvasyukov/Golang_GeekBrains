package main

import (
	"fmt"
)

func Worker(ch chan<- struct{}) {
	ch <- struct{}{}
}

func WorkerHandler(workerQuantity int, handler func()) {
	jobCh := make(chan struct{})
	limit := 0

	for w := 0; w < workerQuantity; w++ {
		go Worker(jobCh)
	}
	for {
		select {
		case <-jobCh:
			handler()
			limit++
			if limit == workerQuantity {
				close(jobCh)
				return
			}
		}
	}
}

func main() {
	count := 0
	handler := func() {
		count++
		fmt.Println(count)
	}
	WorkerHandler(1000, handler)
}
