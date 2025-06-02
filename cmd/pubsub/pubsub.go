package pubsub

import (
	"context"
	"sync"
)

type PubSub interface {
	Publish(ctx context.Context, msg string, topic string)
	Subscruber()
}

type pubSub struct {
	mu   sync.Mutex
	logs map[string][]string // m m m m m
	// topic -> subID
	offsets map[string]map[string]int
}

func New() *pubSub {
	p := &pubSub{
		logs:    make(map[string][]string),
		offsets: make(map[string]map[string]int),
	}

	return p
}

func (pu *pubSub) Publish(ctx context.Context, msg string, topic string) {
	pu.mu.Lock()
	defer pu.mu.Unlock()

	_, ok := pu.offsets[topic]
	if !ok {
		pu.offsets[topic] = make(map[string]int)
	}

	pu.logs[topic] = append(pu.logs[topic], msg)
}

func (pu *pubSub) Subscruber(id string) *sub {
	return &sub{
		pubSub: pu,
		ID:     id,
	}
}
