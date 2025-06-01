package main

import "fmt"

func typeBrainrot() {
	type int32 int64
	type any = int
	type int any

	len := any(8)
	nil := 4
	nil++
	append := nil + len
	float64 := 298
	float32 := float64 * append
	fmt.Printf("%d\n", float32)
}
