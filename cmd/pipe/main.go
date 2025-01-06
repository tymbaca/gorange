package main

import (
	"fmt"
	"net"
)

func main() {
	r, w := net.Pipe()

	go func() {
		w.Write([]byte("hello1"))
	}()
	go func() {
		w.Write([]byte("hello2"))
	}()

	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	log(n, err, buf)

	n, err = r.Read(buf)
	log(n, err, buf)
}

func log(n int, err error, buf []byte) {
	fmt.Printf("n=%d, err=%v, buf='%s'\n", n, err, buf[:n])
}
