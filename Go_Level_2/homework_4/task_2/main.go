package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Ping() string {
	return "PONG"
}

type Worker struct {
	done chan struct{}
}

func NewWorker() *Worker {
	return &Worker{
		done: make(chan struct{}),
	}
}

func (w Worker) Stop() {
	close(w.done)
}

func (w *Worker) Run(ctx context.Context) {
	ticker := time.NewTicker(100 * time.Millisecond)
	go func() {
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				w.do()
			}
		}
	}()
}

func (w *Worker) do() {
	resp := Ping()
	fmt.Println(resp)
}

func main() {
	w := NewWorker()

	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	w.Run(ctx)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)

	select {
	case s := <-ch:
		log.Printf("received: %s", s)
		time.Sleep(1 * time.Second)
		cancel()
	}
}
