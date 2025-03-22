package main

import (
	"context"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/charmbracelet/log"
	cacherpc "github.com/tymbaca/gorange/cmd/election/cache/rpc"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	c := cacherpc.Connect("localhost:1234")
	id := strconv.Itoa(rand.Intn(100))
	log.Print("starting", "my-id", id)

	for {
		if chance(0) {
			sleep := time.Duration(rand.Intn(10000)) * time.Millisecond
			log.Error("shotdown! crit-%^ic.. l ER0-OR...", "wake-after", sleep)

			select {
			case <-ctx.Done():
				log.Fatal("exiting..")
				return
			case <-time.After(sleep):
			}
		}

		if ctx.Err() != nil {
			log.Fatal("exiting..")
			return
		}

		time.Sleep(time.Second + time.Duration(rand.Intn(100))*time.Millisecond)
		current, set, err := c.SetNX("leader", id, 5*time.Second)
		if err != nil {
			panic(err)
		}

		if set {
			log.Info("setnx", "current", current, "set", set)
		} else {
			log.Print("setnx", "current", current, "set", set)
		}
	}
}

func chance(percent float64) bool {
	dice := rand.Float64()

	return dice < percent/100
}
