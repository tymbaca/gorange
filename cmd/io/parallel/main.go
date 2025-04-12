package main

import (
	"errors"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	in, out := net.Pipe()

	go func() {
		for range 100 {
			_, err := in.Write([]byte("1"))
			if err != nil {
				panic(err)
			}

			time.Sleep(10 * time.Millisecond)
		}
	}()
	go func() {
		for range 100 {
			_, err := in.Write([]byte("2"))
			if err != nil {
				panic(err)
			}
			time.Sleep(10 * time.Millisecond)
		}
	}()

	buf := make([]byte, 1024)
	for {
		n, err := out.Read(buf)
		if n > 0 {
			os.Stdout.Write(buf[:n])
		}
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			panic(err)
		}

	}
}
