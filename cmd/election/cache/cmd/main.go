package main

import (
	"log"

	"github.com/tymbaca/gorange/cmd/election/cache"
	cacherpc "github.com/tymbaca/gorange/cmd/election/cache/rpc"
)

func main() {
	c := cache.New()
	port := ":1234"
	log.Println("serving on", port)
	log.Fatal(cacherpc.Serve(c, port))
}
