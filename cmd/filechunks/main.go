package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

type Chunk struct {
	ID   int
	Size uint64
	Body io.Reader
}

func writeChunkConcurrently(chunks []*Chunk, outputFile string) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	offset := int64(0)

	// Start goroutines for each chunk
	for _, chunk := range chunks {
		wg.Add(1)
		go func(c *Chunk) {
			defer wg.Done()
			// Create the output file
			f, err := os.Create(outputFile)
			if err != nil {
				panic(err)
			}
			defer f.Close()

			// Acquire the mutex to ensure atomic offset updates
			mu.Lock()
			chunkOffset := offset
			offset += int64(c.Size)

			// Seek to the correct offset in the output file
			_, err = f.Seek(chunkOffset, io.SeekStart)
			if err != nil {
				fmt.Printf("Error seeking to offset %d: %v\n", chunkOffset, err)
				return
			}
			mu.Unlock()

			// Write the chunk's contents
			_, err = io.Copy(f, c.Body)
			if err != nil {
				fmt.Printf("Error writing chunk %d: %v\n", c.ID, err)
				return
			}
		}(chunk)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	return nil
}

func main() {
	// Example chunks
	chunks := []*Chunk{
		{ID: 1, Size: 1024, Body: strings.NewReader("Chunk 1 content")},
		{ID: 2, Size: 2048, Body: strings.NewReader("Chunk 2 content")},
		{ID: 3, Size: 512, Body: strings.NewReader("Chunk 3 content")},
	}

	err := writeChunkConcurrently(chunks, "cmd/filechunks/output.txt")
	if err != nil {
		fmt.Println("Error writing chunks:", err)
	} else {
		fmt.Println("Chunks written concurrently to output.txt")
	}
}
