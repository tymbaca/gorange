package main

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	cacherpc "github.com/tymbaca/gorange/cmd/election/cache/rpc"
)

func main() {
	c := cacherpc.Connect("localhost:1234")
	id := strconv.Itoa(rand.Intn(100))
	log.Println("my id is:", id)

	for {
		time.Sleep(time.Second)
		current, ok, err := c.SetNX("leader", id, 5*time.Second)
		if err != nil {
			panic(err)
		}

		log.Println("current:", current, "set:", ok)
	}
}
