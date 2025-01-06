package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type ChunkReader struct {
	file   *os.File
	offset int64
	limit  int64
	mu     sync.Mutex
}

func NewChunkReader(file *os.File, offset, size int64) *ChunkReader {
	return &ChunkReader{
		file:   file,
		offset: offset,
		limit:  offset + size,
	}
}

func (cr *ChunkReader) Read(p []byte) (n int, err error) {
	cr.mu.Lock()
	defer cr.mu.Unlock()

	// Check if we've reached the end of the chunk
	if cr.offset >= cr.limit {
		return 0, io.EOF
	}

	if cr.offset+int64(len(p)) >= cr.limit {
		p = p[:(cr.limit - cr.offset)]
	}

	// Read from the file, limiting the read to the chunk size
	n, err = cr.file.ReadAt(p, cr.offset)
	if err != nil {
		return
	}

	// Update the offset for the next read
	cr.offset += int64(n)

	return
}

func main() {
	// Open the file
	f, err := os.Open("cmd/multireader/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Create chunk readers for different parts of the file
	chunk1 := NewChunkReader(f, 0, 10)
	chunk2 := NewChunkReader(f, 10, 10)
	chunk3 := NewChunkReader(f, 20, 10)
	_, _, _ = chunk1, chunk2, chunk3

	// Read from the chunks on demand
	// data1, err := io.ReadAll(chunk1)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Chunk 1:", string(data1))
	//
	// data2, err := io.ReadAll(chunk2)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Chunk 2:", string(data2))
	//
	// data3, err := io.ReadAll(chunk3)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Chunk 3:", string(data3))

	full := io.MultiReader(chunk1, chunk2, chunk3)

	buf := make([]byte, 3)
	fmt.Print("Full: ")
	for {
		time.Sleep(100 * time.Millisecond)
		n, err := full.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Print("\n")
				return
			}

			panic(err)
		}
		fmt.Print(string(buf[:n]))
		// fmt.Println(n)
	}

}

//
// func main() {
// 	r1 := strings.NewReader("this is first chunk ")
// 	r2 := strings.NewReader("this is second chunk ")
// 	r3 := strings.NewReader("this is third chunk ")
//
// 	r := io.MultiReader(r1, r2, r3)
//
// 	buf := make([]byte, 5)
// 	for {
// 		time.Sleep(100 * time.Millisecond)
// 		n, err := r.Read(buf)
// 		if err != nil {
// 			if errors.Is(err, io.EOF) {
// 				fmt.Println("good")
// 				return
// 			}
//
// 			panic(err)
// 		}
//
// 		fmt.Print(string(buf[:n]))
// 	}
// }
