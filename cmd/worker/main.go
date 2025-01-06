package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// ctx := context.Background()
	// svc := &service{}
	//
	// wrk := worker.New(svc.DeleteBadUsers, 10*time.Second)
	//
	// wrk.Run(ctx)

	ch := make(chan int)
	close(ch)
	// var ch chan int

	select {
	case i, ok := <-ch:
		fmt.Println(i, ok)
	default:
	}

	fmt.Println("exit")
}

type service struct{}

func (s *service) DeleteBadUsers(ctx context.Context) error {
	time.Sleep(2 * time.Second)

	if rand.Int()%2 == 0 { // 124151 % 2 != 0 | 53250 % 2 == 0
		return nil
	} else {
		return errors.New("some error")
	}
}
