package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"time"
	"unsafe"
)

type Fileds []string

func main() {
	printBits(-1)
	printBits(0)
	printBits(1)
	v127 := int8(127)
	v127++
	printBits(v127)
}

func printBits(v int8) {
	byte := *(*byte)(unsafe.Pointer(&v))
	fmt.Printf("%08b\n", byte)
}

func acceptStrings(args ...string) {
}

func client() {
	time.Sleep(1 * time.Second)
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	conn.Write([]byte("hello"))
}

func server() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	conn, err := l.Accept()
	if err != nil {
		panic(err)
	}
	conn.SetReadDeadline(time.Now().Add(3 * time.Second))

	buf := make([]byte, 5)
	_, err = conn.Read(buf)
	if err != nil {
		panic(err)
	}

	// block forever
	_, err = conn.Read(buf)
	if errors.Is(err, os.ErrDeadlineExceeded) {
		fmt.Println("i/o deadline!!!")
		return
	}
	if err != nil {
		panic(err)
	}
}
