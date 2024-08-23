package cache

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"sync/atomic"
	"time"

	"github.com/charmbracelet/log"
)

type DB[K comparable, V any] interface {
	Get(key K) (V, error)
	MGet(keys []K) ([]*V, error)
	Keys() ([]K, error)
}

type Cache[K comparable, V any] struct {
	db DB[K, V]

	mu     sync.RWMutex
	data   map[K]item[V]
	ttl    time.Duration
	jitter time.Duration

	hits          atomic.Uint64
	misses        atomic.Uint64
	invalidateDur time.Duration
}

type item[V any] struct {
	Val      V
	Deadline time.Time
}

func NewCache[K comparable, V any](ctx context.Context, db DB[K, V], opts ...Option) *Cache[K, V] {
	cache := &Cache[K, V]{
		db:   db,
		ttl:  10 * time.Second,
		data: make(map[K]item[V]),
	}
	for _, o := range opts {
		o(cache)
	}

	cache.jitter = cache.ttl / 4

	go cache.startInvalidation(ctx)

	return cache
}

func (c *Cache[K, V]) Get(key K) (V, error) {
	c.mu.RLock()
	itm, ok := c.data[key]
	c.mu.RUnlock()
	if ok {
		c.hits.Add(1)
		return itm.Val, nil
	}
	c.misses.Add(1)

	val, err := c.db.Get(key)
	if err != nil {
		return val, err
	}

	jitter := randDur(-c.jitter, c.jitter)
	itm = item[V]{Val: val, Deadline: time.Now().Add(c.ttl + jitter)}

	c.mu.Lock()
	c.data[key] = itm
	c.mu.Unlock()

	log.Debugf("adding item: %+v", itm)

	return val, nil
}

func randDur(min, max time.Duration) time.Duration {
	return time.Duration(rand.Int64N(int64(max-min))) + min
}

func (c *Cache[K, V]) MGet(keys []K) ([]*V, error) {
	return c.db.MGet(keys)
}
func (c *Cache[K, V]) Keys() ([]K, error) {
	return c.db.Keys()
}

func (c *Cache[_, _]) Stats() string {
	return fmt.Sprintf("hits: %d, misses: %d, data len: %d, invalidate dur: %s", c.hits.Load(), c.misses.Load(), len(c.data), c.invalidateDur)
}

func (c *Cache[_, _]) startInvalidation(ctx context.Context) {
	tick := time.NewTicker(10 * time.Second)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			c.invalidate()
		}
	}
}

func (c *Cache[K, V]) invalidate() {
	now := time.Now()
	newData := make(map[K]item[V], len(c.data))

	c.mu.RLock()
	for key, item := range c.data {
		if !item.Deadline.Before(now) {
			newData[key] = item
		}
	}
	c.mu.RUnlock()

	c.mu.Lock()
	c.data = newData
	c.mu.Unlock()

	c.invalidateDur = time.Since(now)
}
