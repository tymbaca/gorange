package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("beep", i)
		i++
	}
}
