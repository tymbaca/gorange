package cache

import "time"

type Option = func(c *Cache)

func WithTTL(dur time.Duration) Option {
	return func(c *Cache) {
		c.ttl = dur
	}
}

func WithPrealloc(size int) Option {
	return func(c *Cache) {
		c.data = make(map[string]item, size)
	}
}
