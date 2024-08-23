package cache

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func BenchmarkMapInvalidate(b *testing.B) {
	m := make(map[string]item)
	for i := range 100_000 {
		key := strconv.Itoa(i)
		deadline := time.Now().Add(1 * time.Hour)
		if i%200 == 0 {
			deadline = time.Now()
		}

		m[key] = item{
			Val:      key,
			Deadline: deadline,
		}
	}

	b.Run("delete in-place", func(b *testing.B) {
		for range b.N {
			now := time.Now()
			for k, v := range m {
				if v.Deadline.Before(now) {
					delete(m, k)
				}
			}
			fmt.Printf("elapsed %s\n", time.Since(now).String())
		}
	})

	b.Run("copy and filter", func(b *testing.B) {
		for range b.N {
			now := time.Now()
			newM := make(map[string]item, len(m))
			for k, v := range m {
				if !v.Deadline.Before(now) {
					newM[k] = v
				}
			}
			fmt.Printf("elapsed %s\n", time.Since(now).String())
		}
	})
}
