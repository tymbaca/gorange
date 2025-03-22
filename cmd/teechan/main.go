package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	in := make(chan int)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			in <- 10
		}
	}()

	for v := range orDone(ctx, in) {
		time.Sleep(3 * time.Second)
		fmt.Println(v)
	}
}

func orDone(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case val, ok := <-in:
				if !ok {
					return
				}

				// out <- val
				select {
				case <-ctx.Done():
					return
				case out <- val:
				}
			}
		}
	}()

	return out
}
