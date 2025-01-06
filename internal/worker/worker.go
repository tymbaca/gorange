package worker

import (
	"context"
	"log"
	"time"
)

/*
func()

every 5 sec

Run
*/

type Exec func(ctx context.Context) error

type Worker struct {
	exec     Exec
	interval time.Duration
}

func New(exec Exec, interval time.Duration) *Worker {
	return &Worker{
		exec:     exec,
		interval: interval,
	}
}

func (w *Worker) Run(ctx context.Context) {
	go func() {
		for {
			err := w.exec(ctx)
			log.Println("start")
			if err != nil {
				log.Println(err.Error())
			}
			log.Println("end")

			time.Sleep(w.interval)
		}
	}()
}

/*
main  :  -------- Run go func()
func():                    {   }---- for exec ....
*/

func (w *Worker) RunSilent(ctx context.Context) {
	go func() {
		for {
			err := w.exec(ctx)
			if err != nil {
				log.Println(err.Error())
			}

			time.Sleep(w.interval)
		}
	}()
}
