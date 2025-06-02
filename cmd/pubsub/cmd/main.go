package main

import (
	"context"
	"log"
	"time"

	"github.com/tymbaca/gorange/cmd/pubsub"
)

const topic = "topic1"

func main() {
	ctx := context.Background()

	ps := pubsub.New()

	ps.Publish(ctx, "1", topic)
	ps.Publish(ctx, "2", topic)
	ps.Publish(ctx, "3", topic)

	go func() {
		sub := ps.Subscruber("id1")

		for {
			msg, err := sub.Listen(ctx, topic)
			if err != nil {
				panic(err)
			}

			time.Sleep(1 * time.Second)
			log.Printf("sub id %s, got msg %s", sub.ID, msg)
		}
	}()

	go func() {
		sub := ps.Subscruber("id2")

		for {
			msg, err := sub.Listen(ctx, topic)
			if err != nil {
				panic(err)
			}

			time.Sleep(2 * time.Second)
			log.Printf("sub id %s, got msg %s", sub.ID, msg)
		}
	}()

	go func() {
		sub := ps.Subscruber("id3")

		for {
			msg, err := sub.Listen(ctx, topic)
			if err != nil {
				panic(err)
			}

			time.Sleep(3 * time.Second)
			log.Printf("sub id %s, got msg %s", sub.ID, msg)
		}
	}()

	time.Sleep(1 * time.Minute)
}
