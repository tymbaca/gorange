package main

import (
	"context"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	_ = ctx

	cancel()
	cancel()
	cancel()
}
