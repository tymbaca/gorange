package main

import "github.com/tymbaca/gorange/internal/interface/somelib"

func main() {
	m := myDoer{}
	somelib.UseDoer(m, 10)
}

type myDoer struct {
}

func (d myDoer) do(val int) (int, error) {
	return val * val, nil
}
