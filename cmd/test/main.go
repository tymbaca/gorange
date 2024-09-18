package main

import "fmt"

func main() {
	i, err := f()
	fmt.Println(&i, &err)

	a, err := f()
	fmt.Println(&a, &err)

	a, err = f()
	fmt.Println(&a, &err)
}

func f() (int, error) {
	return 129, nil
}
