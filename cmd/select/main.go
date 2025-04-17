package main

import "fmt"

func main() {
	ch := make(chan int, 1000)

	select {
	case ch <- w(f1()):
	case ch <- f2():
	case ch <- f3():
	}
}

func w(i int) int {
	fmt.Println("w")
	return i
}

func f1() int {
	fmt.Println("f1")
	return 1
}

func f2() int {
	fmt.Println("f2")
	return 2
}

func f3() int {
	fmt.Println("f3")
	return 3
}
