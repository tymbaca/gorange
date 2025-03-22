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

	// TODO: retry connection
	c := cacherpc.Connect("localhost:1234")
	id := strconv.Itoa(rand.Intn(100))
	log.Print("starting", "my-id", id)

	for {
		maybeCrush(ctx, 0)

		if ctx.Err() != nil {
			log.Fatal("exiting..")
			return
		}

		time.Sleep(time.Second + time.Duration(rand.Intn(100))*time.Millisecond)
		current, set, err := c.SetNX("leader", id, 5*time.Second)
		if err != nil {
			log.Error("can't setnx", "err", err)
		}

		if set {
			log.Info("setnx", "current", current, "set", set)
			hold(ctx, c, id)
		} else {
			log.Print("setnx", "current", current, "set", set)
		}
	}
}

const holdpref = "====HOLD==== "

func hold(ctx context.Context, cache *cacherpc.Client, id string) {
	log.Info(holdpref+"we are leader now", "id", id)
	for {
		maybeCrush(ctx, 15)

		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Second + time.Duration(rand.Intn(100))*time.Millisecond):
		}

		current, set, err := cache.Prolong("leader", 5*time.Second)
		if err != nil {
			log.Error(holdpref+"can't prolong", "err", err)
			continue
		}
		log.Info(holdpref + "prolonged")

		if current != id && set {
			log.Warn(holdpref+"we lost leadership, following..", "current", current)
			return
		}
	}
}

func maybeCrush(ctx context.Context, percent float64) {
	if chance(percent) {
		sleep := time.Duration(rand.Intn(10000))*time.Millisecond + 5*time.Second
		log.Error("shotdown! crit-%^ic.. l ER0-OR...", "wake-after", sleep)

		select {
		case <-ctx.Done():
			log.Fatal("exiting..")
			return
		case <-time.After(sleep):
		}
	}
}

func chance(percent float64) bool {
	dice := rand.Float64()

	return dice < percent/100
}
