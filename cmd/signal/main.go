package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"gitlab.wildberries.ru/logisticcloud/platform/lib/logger"
	"gitlab.wildberries.ru/logisticcloud/platform/lib/task-manager/job"
)

func main() {
	logger.Init("local")
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	job.NewInterval(job.WithDetachCtx()).Start(ctx, "some-work", 1*time.Second, doSomeWork)
	// job.NewInterval().Start(ctx, "some-work", 1*time.Second, doSomeWork)

	<-ctx.Done()
	fmt.Printf("exiting: %v\n", ctx.Err())
}

func doSomeWork(ctx context.Context) error {
	for i := range 10 {
		time.Sleep(100 * time.Millisecond)

		// Якобы делаем сетевой вызов
		if err := ctx.Err(); err != nil {
			logger.Infof("%d\t | shit, you broke the request with your shutdown", i+1)

			return err
		}

		logger.Infof("%d\t | successful request", i+1)
	}

	return nil
}
