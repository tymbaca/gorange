package pubsub

import (
	"context"
	"sync"

	"github.com/google/uuid"
)

type PubSub interface {
	Publish(ctx context.Context, msg string, topic string)
	Subscribe() Sub // uuid
}

func New(ctx context.Context) *pubSub {
	p := &pubSub{
		logs:          make(map[string][]string),
		offsets:       make(map[string]map[string]int),
		offsetOffsets: make(map[string]int),
	}

	go p.runCleaner(ctx)

	return p
}

type pubSub struct {
	globalMu      sync.RWMutex
	offsetOffsets map[string]int

	logMu sync.RWMutex
	logs  map[string][]string

	offsetsMu sync.Mutex
	// topic -> id
	offsets map[string]map[string]int
}

func (p *pubSub) Publish(ctx context.Context, msg string, topic string) {
	p.globalMu.RLock()
	defer p.globalMu.RUnlock()

	p.logMu.Lock()
	defer p.logMu.Unlock()

	p.offsetsMu.Lock()
	defer p.offsetsMu.Unlock()

	p.logs[topic] = append(p.logs[topic], msg)
	p.offsets[topic] = make(map[string]int)
}

func (p *pubSub) Subscribe() Sub {
	p.globalMu.RLock()
	defer p.globalMu.RUnlock()

	p.logMu.RLock()
	defer p.logMu.RUnlock()

	p.offsetsMu.Lock()
	defer p.offsetsMu.Unlock()

	id := uuid.NewString()

	for topic := range p.logs {
		p.offsets[topic][id] = p.offsetOffsets[topic] // WARN: can go wrong
	}

	return &sub{
		pubSub: p,
		id:     id,
	}
}

func (p *pubSub) runCleaner(ctx context.Context) {
	// for {
	// 	select {
	// 	case <-ctx.Done():
	// 		return
	// 	case <-time.After(5 * time.Second):
	// 		if needToClean(p.offsets, len(p.log)) {
	// 			p.clean()
	// 		}
	// 	}
	// }
}

func needToClean(offsets map[string]int, logSize int) bool {
	panic("not implemented")
}

func (p *pubSub) clean() {
	p.globalMu.Lock()
	defer p.globalMu.Unlock()

	panic("not implemented")
}
