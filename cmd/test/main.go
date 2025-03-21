package main

import (
	"fmt"
	"time"
)

type Hash [32]byte

func main() {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		i++
		fmt.Print("\033[H\033[2J")
		fmt.Println(i)
	}

	// h1 := Hash{}
	// h2 := Hash{}
	//
	// fmt.Println(h1 == h2)
}

func sum(a, b int) int {
	return a + b
}
