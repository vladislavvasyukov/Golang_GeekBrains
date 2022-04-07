package main

import "fmt"

func Worker(ch chan<- struct{}) {
	ch <- struct{}{}
}

func WorkerHandler(workerQuantity int, handler func()) {
	// Запускаются воркеры в указанном количестве. Каждый выполняет свою работу (handler) в переданный ему канал и
	// заверщает свою работу. Когда количество выполненных воркеров достигнет лимита, канал закроется и произойдет
	// выход из функции
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
	// 1. Задание
	// С помощью пула воркеров написать программу, которая запускает 1000 горутин, каждая из
	// которых увеличивает число на 1. Дождаться завершения всех горутин и убедиться, что при
	// каждом запуске программы итоговое число равно 1000.

	count := 0
	handler := func() {
		count++
		fmt.Println(count)
	}
	WorkerHandler(1000, handler)

	// 2. Задание
	// Написать программу, которая при получении в канал сигнала SIGTERM останавливается не
	// позднее, чем за одну секунду (установить таймаут).
}
