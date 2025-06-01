package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/tymbaca/gorange/cmd/pubsub/pubsub"
)

// queue
// buffering
// no block

const topic = "topic1"

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	ps := pubsub.New(ctx)

	go func() {
		i := 1
		for range time.Tick(500 * time.Millisecond) {
			ps.Publish(ctx, strconv.Itoa(i), topic)
			i++
		}
	}()

	go func() {
		s := ps.Subscribe()

		for {
			msg, err := s.Listen(ctx, topic)
			if err != nil {
				log.Println(err)
				time.Sleep(1 * time.Second)
				continue
			}

			log.Println("msg", msg)
		}
	}()

	<-ctx.Done()
}
