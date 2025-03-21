package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/tymbaca/gorange/cmd/rpc/server"
)

func main() {
	arith := new(server.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	log.Println("ready")
	log.Fatal(http.Serve(l, nil))
}
