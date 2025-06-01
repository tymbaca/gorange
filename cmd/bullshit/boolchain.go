package bullshit

import (
	"context"
	"fmt"
	"time"
)

func boonChain() {
	ctx := context.Background()
	if check1(ctx) && check2(ctx) && check3(ctx) {
		fmt.Println("all three are good")
	}
}

func check1(ctx context.Context) bool {
	time.Sleep(1 * time.Second)
	return true
}

func check2(ctx context.Context) bool {
	time.Sleep(1 * time.Second)
	return false
}

func check3(ctx context.Context) bool {
	time.Sleep(1 * time.Second)
	return true
}
