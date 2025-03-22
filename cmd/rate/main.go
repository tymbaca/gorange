package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	r := rate.NewLimiter(rate.Limit(0.16666), 10)

	i := 0
	for {
		i++
		start := time.Now()

		err := r.Wait(ctx)
		if err != nil {
			panic(err)
		}

		log.Println("req", i, time.Since(start))
	}
}
