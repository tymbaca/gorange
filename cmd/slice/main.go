package main

import "fmt"

type Test struct {
	a int16
	b int8
	c int8
	A int64
}

type TestPacked struct {
	// A int64
	a int16
	b int8
	c int8
}

func main() {
	a := make([]int, 0, 5)

	foo(a)

	a[:5][4] = 7
	fmt.Println(a[:5])

	// s := make([]byte, 10)
	// fmt.Printf("len=%d, cap=%d, val=%#v\n", len(s), cap(s), s)
	//
	// s = s[:5]
	// fmt.Printf("len=%d, cap=%d, val=%#v\n", len(s), cap(s), s)
}

func foo(a []int) {
	a = append(a, 10)
}

var _count = 5

func decr() (int, bool) {
	if _count == 0 {
		return 0, false
	}

	val := _count
	_count--
	return val, true
}
