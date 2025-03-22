package main

import (
	"fmt"
	"log"
)

func main() {
	var quacker Quacker
	duck := GetDuck()
	if duck != nil {
		quacker = duck
	}

	checkQuacker(quacker)
}

func checkQuacker(quacker Quacker) {
	if quacker == nil {
		log.Fatalf("it's nil, %#v", quacker)
	}

	log.Fatalf("it's not nil, %#v", quacker)
}

type Quacker interface {
	Quack() string
}

type Duck struct {
	Name string
}

func (d *Duck) Quack() string {
	return fmt.Sprintf("QUACK, I'M %s", d.Name)
}

func GetDuck() *Duck {
	return nil
}
