package test

import "fmt"

type Hash [32]byte

func main() {
	h1 := Hash{}
	h2 := Hash{}

	fmt.Println(h1 == h2)
}

func sum(a, b int) int {
	return a + b
}
