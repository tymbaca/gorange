package main

import (
	"fmt"
)

var s = []int{
	10: 10,
	11,
	12,
	5: 5,
	7: 7,
	8,
}

func main() {
	for i, val := range s {
		if val == 0 {
			fmt.Println(i, "ZERO")
		} else {
			fmt.Println(i, val)
		}
	}
}
