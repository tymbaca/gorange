package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand/v2"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/charmbracelet/log"
	"github.com/tymbaca/gorange/internal/cache"
	"github.com/tymbaca/gorange/internal/helper/mem"
	"golang.org/x/sync/errgroup"
)

const (
	N        = 100_000
	_ttl     = 1 * time.Minute
	_rps     = 300
	_latency = 30 * time.Millisecond
)

func main() {
	log.SetLevel(log.DebugLevel)

	data := make(map[string]string, N)
	for i := range N {
		key := strconv.Itoa(i)
		data[key] = "val of " + key
	}

	db := DB{data}
	cache := cache.NewCache(
		context.Background(),
		db,
		cache.WithTTL(_ttl),
		cache.WithPrealloc(N),
	)

	client := &Client{
		cache: cache,
	}
	go client.updateLatency()

	g := errgroup.Group{}
	for id := range 10 {
		g.Go(client.runGetter(id))
	}

	http.HandleFunc("/cache/stats", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(cache.Stats()))
	})
	http.HandleFunc("/client/stats", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(client.Stats()))
	})
	http.HandleFunc("/mem/stats", func(w http.ResponseWriter, r *http.Request) {
		stats := mem.FormatMem(mem.KiB)
		w.Write([]byte(stats))
	})
	http.ListenAndServe(":8000", nil)

	g.Wait()
}

type Client[K comparable, V any] struct {
	cache *cache.Cache[K, V]

	latency    time.Duration
	totalCalls atomic.Int64
	totalDur   atomic.Int64
}

func (c *Client[K, V]) runGetter(id int) func() error {
	return func() error {
		for {
			select {
			case <-time.Tick(time.Second / time.Duration(_rps)):
				log.Infof("%d goroutine: get", id)
				start := time.Now()

				key := strconv.Itoa(rand.IntN(N))
				_, err := c.cache.Get(key)
				if err != nil {
					log.Errorf("%d goroutine: error from get: %s", id, err.Error())
				}

				c.totalCalls.Add(1)
				c.totalDur.Add(int64(time.Since(start)))
			}
		}
	}
}

func (c *Client) updateLatency() {
	for {
		select {
		case <-time.Tick(3 * time.Second):
			if c.totalCalls.Load() == 0 {
				continue
			}

			c.latency = time.Duration(c.totalDur.Load() / c.totalCalls.Load())

			c.totalCalls.Store(0)
			c.totalDur.Store(0)
		}
	}
}

func (c *Client) Stats() string {
	return fmt.Sprintf("averate latency: %s, total calls: %d", c.latency.String(), c.totalCalls.Load())
}

type DB struct {
	data map[string]string
}

func (d DB) Get(key string) (string, error) {
	time.Sleep(_latency)
	val, ok := d.data[key]
	if !ok {
		return "", errors.New("not found")
	}

	return val, nil
}

func (d DB) MGet(keys []string) ([]*string, error) {
	return nil, errors.New("not implemented")
}
func (d DB) Keys() ([]string, error) {
	return nil, errors.New("not implemented")
}
