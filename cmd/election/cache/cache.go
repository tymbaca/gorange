package cache

import (
	"sync"
	"time"

	"github.com/charmbracelet/log"
)

type Cache struct {
	data map[string]record
	mu   sync.RWMutex
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

func (c *Cache) Prolong(key string, ttl time.Duration) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	record, ok := c.data[key]
	if !ok {
		return "", false
	}

	record.Deadline = time.Now().Add(ttl)
	c.data[key] = record

	log.Print("prolong", "key", key, "by", ttl)
	return record.Val, true
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	record, ok := c.data[key]
	if !ok {
		return "", false
	}

	if record.Deadline.Before(time.Now()) {
		return "", false
	}

	return record.Val, ok
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
		c.data[key] = record{Val: val, Deadline: time.Now().Add(ttl)}
		log.Info("setnx", "key", key, "to", val, "by", ttl)
		return val, true
	}

	return current.Val, false
}
