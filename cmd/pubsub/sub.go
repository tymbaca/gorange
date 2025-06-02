package pubsub

import (
	"context"
	"time"
)

type Sub interface {
	Listen(ctx context.Context, topic string) (string, error)
}

type sub struct {
	pubSub *pubSub
	ID     string
}

func (su *sub) Listen(ctx context.Context, topic string) (string, error) {
	if su.canRead(topic) {
		return su.readNext(topic), nil
	}

	if err := su.blockUntilNewMsg(ctx, topic); err != nil {
		return "", err
	}

	return su.readNext(topic), nil
}

func (su *sub) canRead(topic string) bool {
	su.pubSub.mu.Lock()
	defer su.pubSub.mu.Unlock()

	logSize := len(su.pubSub.logs[topic])
	offset := su.pubSub.offsets[topic][su.ID]

	return offset < logSize
}

func (su *sub) blockUntilNewMsg(ctx context.Context, topic string) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(10 * time.Millisecond):
			if su.canRead(topic) {
				return nil
			}
		}
	}
}

func (su *sub) readNext(topic string) string {
	su.pubSub.mu.Lock()
	defer su.pubSub.mu.Unlock()

	offset := su.pubSub.offsets[topic][su.ID]
	msg := su.pubSub.logs[topic][offset]
	su.pubSub.offsets[topic][su.ID]++

	return msg
}
