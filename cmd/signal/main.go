package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"gitlab.wildberries.ru/logisticcloud/platform/lib/logger"
)

func main() {
	logger.Init("local")
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// job.NewInterval(job.WithDetachCtx()).Start(ctx, "some-work", 1*time.Second, doSomeWork)
	// job.NewInterval().Start(ctx, "some-work", 1*time.Second, doSomeWork)

	go func() {
		err := work(ctx)
		if err != nil {
			logger.Error(err, "worker")
		}
	}()

	<-ctx.Done()
	fmt.Printf("exiting: %v\n", ctx.Err())
}

func work(ctx context.Context) error {
	t := time.NewTicker(1 * time.Second)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			// doSomeWork(ctx)
			doSomeWork(context.Background())
		case <-ctx.Done():
			return fmt.Errorf("exiting worker: %w", ctx.Err())
		}
	}
}

func doSomeWork(ctx context.Context) error {
	for i := range 20 {
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
