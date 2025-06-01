package main

import "fmt"

type Offer struct {
	ID    int
	Attrs *Attrs
}

type Attrs struct {
	Name  string
	Price int
}

func main() {
	o1 := Offer{
		ID: 1,
		Attrs: &Attrs{
			Name:  "111",
			Price: 1111,
		},
	}

	fmt.Printf("o1: %#v\n", o1.Attrs)

	o2 := o1
	attrsCopy := *o1.Attrs
	o2.Attrs = &attrsCopy
	o1.Attrs.Name = "222"
	o1.Attrs.Price = 2222

	fmt.Println("after copy")
	fmt.Printf("o1: %#v\n", o1.Attrs)
	fmt.Printf("o2: %#v\n", o2.Attrs)
}
