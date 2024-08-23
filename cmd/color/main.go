package main

import (
	"fmt"
	"strings"
)

func main() {
	Log("<b>Hello</b>")
}

func Log(text string) {
	text = strings.ReplaceAll(text, "<b>", "\033[1m")
	text = strings.ReplaceAll(text, "</b>", "\033[0m")

	fmt.Println(text)
}
