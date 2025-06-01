package pubsub

import (
	"context"
	"sync"

	"golang.org/x/sync/errgroup"
)

type PubSub struct {
	mu          sync.Mutex
	subsByTopic map[string]map[string]*Sub
}

func New() *PubSub {
}

func (pu *PubSub) Subsribe(ctx context.Context, id, topic string) (*Sub, error) {
}

func (pu *PubSub) unsubsribe(ctx context.Context, id, topic string) {
	pu.mu.Lock()
	defer pu.mu.Unlock()

	sub, ok := pu.getSub(id, topic)
	if !ok {
		return
	}

	sub.once.Do(func() {
		close(sub.ch)
	})

	delete(pu.subsByTopic[topic], id)
}

func (pu *PubSub) Send(ctx context.Context, topic string, msg string) error {
	pu.mu.Lock()
	defer pu.mu.Unlock()

	subs, ok := pu.subsByTopic[topic]
	if !ok {
		return nil // TODO:
	}

	var wg errgroup.Group
	for _, sub := range subs {
		wg.Go(func() error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case sub.ch <- msg:
				return nil
			}
		})
	}

	return wg.Wait()
}

func (pu *PubSub) getSub(id, topic string) (*Sub, bool) {
	subs, ok := pu.subsByTopic[topic]
	if !ok {
		return nil, false
	}

	sub, ok := subs[id]
	if !ok {
		return nil, false
	}

	return sub, true
}

func (pu *PubSub) Close() {
	pu.mu.Lock()
	defer pu.mu.Unlock()

	for _, subs := range pu.subsByTopic {
		for _, sub := range subs {
			sub.once.Do(func() {
				close(sub.ch)
			})
		}
	}
}

type Sub struct {
	pubSub *PubSub
	ch     chan string
	once   sync.Once
}

func (s *Sub) Listen(ctx context.Context) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case msg, ok <- s.ch:
		return msg, nil
	}
}

func (s *Sub) Close(ctx context.Context) error {
	// s.pubSub.
}
