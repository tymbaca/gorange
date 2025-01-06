package main

import (
	"errors"
	"fmt"
)

type mytype struct {
	name string
}

func (mt *mytype) GetName() string {
	fmt.Println("hello from method")

	if mt != nil {
		return mt.name
	}

	return ""
}

func main() {
	var err interface{}
	m := make(map[error]int)

	m[err] = 5

	err = errors.New("shit")
	m[err] = 8

	fmt.Println(m)
}
