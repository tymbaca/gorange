package cache

import (
	"sync"
	"time"

	"github.com/charmbracelet/log"
)

type Cache struct {
	data map[string]record
	mu   sync.Mutex
}

func New() *Cache {
	return &Cache{
		data: make(map[string]record),
	}
}

type record struct {
	Val      string
	Deadline time.Time
}

func (c *Cache) Prolong(key string, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	record, ok := c.data[key]
	if !ok {
		return
	}

	record.Deadline = time.Now().Add(ttl)
	c.data[key] = record

	log.Print("prolong:", key, "by", ttl)
}

func (c *Cache) SetNX(key, val string, ttl time.Duration) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	current, ok := c.data[key]
	if !ok {
		c.data[key] = record{Val: val, Deadline: time.Now().Add(ttl)}
		log.Info("setnx", "key", key, "to", val, "by", ttl)
		return val, true
	}

	if current.Deadline.Before(time.Now()) {
		log.Warn("expired", "key", key)

		c.data[key] = record{Val: val, Deadline: time.Now().Add(ttl)}
		log.Info("setnx", "key", key, "to", val, "by", ttl)
		return val, true
	}

	return current.Val, false
}
