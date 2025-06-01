package pubsub

import (
	"context"
	"fmt"
	"time"
)

type Sub interface {
	Listen(ctx context.Context, topic string) (string, error)
}

type sub struct {
	pubSub *pubSub
	id     string
}

func (s *sub) Listen(ctx context.Context, topic string) (string, error) {
	s.pubSub.logMu.RLock()
	_, ok := s.pubSub.logs[topic]
	s.pubSub.logMu.RUnlock()
	if !ok {
		return "", fmt.Errorf("there is no topic like %s", topic)
	}

	if s.canRead(topic) {
		return s.readNext(topic), nil
	}

	if err := s.blockUntilCanRead(ctx, topic); err != nil {
		return "", err
	}

	return s.readNext(topic), nil
}

func (s *sub) canRead(topic string) bool {
	s.pubSub.globalMu.RLock()
	defer s.pubSub.globalMu.RUnlock()

	s.pubSub.offsetsMu.Lock()
	offset, ok := s.pubSub.offsets[topic][s.id]
	if !ok {
		offset = s.pubSub.offsetOffsets[topic]
		s.pubSub.offsets[topic][s.id] = offset
	}
	s.pubSub.offsetsMu.Unlock()

	// transform to real offset
	offset -= s.pubSub.offsetOffsets[topic]

	s.pubSub.logMu.RLock()
	logSize := len(s.pubSub.logs[topic]) // 50
	s.pubSub.logMu.RUnlock()

	return offset < logSize
}

func (s *sub) readNext(topic string) string {
	s.pubSub.globalMu.RLock()
	defer s.pubSub.globalMu.RUnlock()

	s.pubSub.offsetsMu.Lock()
	// get the old offset
	offset := s.pubSub.offsets[topic][s.id]
	// increment
	s.pubSub.offsets[topic][s.id]++
	s.pubSub.offsetsMu.Unlock()

	// transform to real offset
	offset -= s.pubSub.offsetOffsets[topic]

	s.pubSub.logMu.RLock()
	msg := s.pubSub.logs[topic][offset]
	s.pubSub.logMu.RUnlock()

	return msg
}

func (s *sub) blockUntilCanRead(ctx context.Context, topic string) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(10 * time.Millisecond):
			if s.canRead(topic) {
				return nil
			}
		}
	}
}
