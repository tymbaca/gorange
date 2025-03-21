package main

import (
	"fmt"
	"net/rpc"

	"github.com/tymbaca/gorange/cmd/rpc/server"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	var reply server.Quotient
	err = client.Call("Arith.Divide", &server.Args{A: 8, B: 2}, &reply)
	if err != nil {
		panic(err)
	}

	fmt.Println("got", reply)
}
