package main

import "fmt"

type Auth struct {
	ClientID string `json:"clientID"`
	ApiKey   string `json:"apiKey"`
}

type MyInt int

func main() {
	{
		var a any = int(7)
		var b any = int(10)
		var c any = int(10)

		fmt.Println(a == b)
		fmt.Println(b == c)
	}

	{
		var a any = []int{1, 2, 3}
		var b any = []int{1, 2, 3, 4}
		var c any = []int{1, 2, 3, 4}

		fmt.Println(a == b)
		fmt.Println(b == c)
	}
	// 	var a any = true
	// 	var b any = 2
	// 	var c any = 3
	//
	// 	if a == b {
	// 		fmt.Println("eq")
	// 	} else {
	// 		fmt.Println("not eq")
	// 	}
	//
	// 	if b == c {
	// 		fmt.Println("eq")
	// 	} else {
	// 		fmt.Println("not eq")
	// 	}
	//
	// 	c = 2
	//
	// 	if b == c {
	// 		fmt.Println("eq")
	// 	} else {
	// 		fmt.Println("not eq")
	// 	}
	//
	// 	c = MyInt(2)
	//
	// 	if b == c {
	// 		fmt.Println("eq")
	// 	} else {
	// 		fmt.Println("not eq")
	// 	}
	// }
	//
	// {
	// 	var a any = []int{1, 2, 3}
	// 	var b any = []int{1, 2, 3}
	//
	// 	if a == b {
	// 		fmt.Println("eq")
	// 	} else {
	// 		fmt.Println("not eq")
	// 	}
	// }
	//
	// {
	// 	var a any = map[int]int{1: 1, 2: 2, 3: 3}
	// 	var b any = map[int]int{1: 1, 2: 2, 3: 3}
	//
	// 	if a == b {
	// 		fmt.Println("eq")
	// 	} else {
	// 		fmt.Println("not eq")
	// 	}
	// }
}
