package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	lis, err := net.Listen("tcp", ":8001")
	if err != nil {
		panic(err)
	}

	fmt.Println("listening")
	for {
		conn, err := lis.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("\naccepted conn")

		_, err = io.Copy(os.Stdout, conn)
		if err != nil {
			panic(err)
		}
	}
}
