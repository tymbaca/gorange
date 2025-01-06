package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	sched := cron.New()
	schedule := cron.Every(5 * time.Second)

	sched.Schedule(schedule, &job{})

	sched.Start()
	select {}
}

type job struct{}

func (j *job) Run() {
	id := rand.Intn(1000)
	fmt.Printf("starting job (id %d)\n", id)
	time.Sleep(20 * time.Second)
	fmt.Printf("ending job (id %d)\n", id)
}
