package main

import (
	"context"
	"fmt"
	"time"
)

type Event struct {
	Val int
	time.Time
}

func WithCallback(ctx context.Context, callback func()) context.Context {
	go func() {
		select {
		case <-ctx.Done():
			callback()
		}
	}()
	return ctx
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	start := time.Now()
	ctx = WithCallback(ctx, func() {
		fmt.Printf("callback called: time since start: %s\n", time.Since(start).String())
	})

	time.Sleep(4 * time.Second)
	fmt.Printf("exiting: time since start: %s\n", time.Since(start).String())
}
