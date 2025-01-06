package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println(runtime.GOMAXPROCS(0))

	for range 100 {
		go func() {
			for {
				// time.Sleep(50 * time.Millisecond)
				// fmt.Print("hello")
			}
		}()
	}

	time.Sleep(50 * time.Second)
}
