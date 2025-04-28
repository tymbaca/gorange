package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	// []int{124, 33, 6, 31, 341, 63, 2}

	var s []*int
	for range 1_0_000_000 {
		s = append(s, ptr(rand.Int()))
	}
	start := time.Now()
	res := maxrec(s, -1)
	fmt.Printf("maxrec: %v, dur: %s\n", res, time.Since(start))

	start = time.Now()
	res = maxloop(s, -1)
	fmt.Printf("maxloop: %v, dur: %s\n", res, time.Since(start))
}

func ptr(n int) *int {
	return &n
}

func maxrec(list []*int, lastMax int) int {
	if len(list) == 0 {
		return lastMax
	}

	elem := list[0]
	// fmt.Printf("%p\n", elem)

	if *elem > lastMax {
		lastMax = *elem
	}
	return maxrec(list[1:], lastMax)
}

func maxloop(list []*int, lastMax int) int {
	for _, n := range list {
		// fmt.Printf("%p\n", n)
		if *n > lastMax {
			lastMax = *n
		}
	}

	return lastMax
}
