package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	for i := range 3 {
		i := i
		go func() {
			for {
				switch i % 3 {
				case 0:
					fmt.Print(" ")
				case 1:
					fmt.Print("|")
				case 2:
					fmt.Print("\n")
				}
				// runtime.Gosched()
			}
		}()
	}

	time.Sleep(10 * time.Second)
}
