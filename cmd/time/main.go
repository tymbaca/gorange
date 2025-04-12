package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Time{}
	fmt.Printf("t1.IsZero(): %v\n", t1.IsZero())
	fmt.Printf("(t1 == (time.Time{})): %v\n", (t1 == (time.Time{})))
}
