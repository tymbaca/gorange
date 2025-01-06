package main

type Test struct {
	a int16
	b int8
	c int8
	A int64
}

type TestPacked struct {
	a int16
	b int8
	A int64
	c int8
}

func main() {
	// s := make([]byte, 10)
	// fmt.Printf("len=%d, cap=%d, val=%#v\n", len(s), cap(s), s)
	//
	// s = s[:5]
	// fmt.Printf("len=%d, cap=%d, val=%#v\n", len(s), cap(s), s)
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
