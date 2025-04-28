package main

import (
	"fmt"
	"time"
)

func main() {
	closeWhenSending()
}

func closeWhenSending() {
	ch := make(chan int)

	go func() {
		<-time.After(150 * time.Millisecond)
		close(ch)
		fmt.Printf("closer: chan closed\n")
	}()

	go func() {
		<-time.After(300 * time.Millisecond)
		for i := range ch {
			fmt.Printf("recv: got %v\n", i)
		}
		fmt.Printf("recv: chan closed\n")
	}()

	ch <- 777
	fmt.Printf("sender: val sent\n")

	select {}
}
