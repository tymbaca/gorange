package circuitbreaker

import (
	"context"
	"fmt"
	"sync"
	"time"

	"google.golang.org/grpc"
)

type UnaryClientInterceptor func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error

func WithCircuitBreaker(cb CircuitBreakerer) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		if cb.IsOpen() {
			return fmt.Errorf("circuitbreaker is open, too many errors, method: %s", method)
		}

		err := invoker(ctx, method, req, reply, cc, opts...)
		cb.Count(err)

		return err
	}
}

type CircuitBreakerer interface {
	IsOpen() bool
	Count(err error)
}

// if last [queueSize] calls have more then [openPercent] percent of errors then open
// while open, wait [cooldown] duration.
// after that switch to half-open.
// while half-open. pass only [passEveryN] element, return error on others
// if got any error in half-open, switch to open
// otherwise, if [successInRowToClose] calls in rows will succeed - switch to closed
type CircuitBreaker struct {
	state state

	queueSize int
	queuePtr  int     // for cyclic queue
	queue     []error // maybe just bool?
	mu        sync.Mutex

	lastOpened          time.Time
	cooldown            time.Duration
	passEveryN          int
	successInRowToClose int
	successInRow        int
}

type state int

const (
	open state = iota
	halfOpen
	closed
)

// maybe cb.Err()?
func (cb *CircuitBreaker) IsOpen() bool {
	if cb.state == open {
		return true
	}

	return false
}

func (cb *CircuitBreaker) Count(err error) {
	switch cb.state {
	case open:
	case halfOpen:
	case closed:
	}
	if len(cb.queue) != cb.queueSize {
		cb.queue = append(cb.queue, err)
		return
	}

	cb.queue[cb.queuePtr] = err

	cb.queuePtr++
	if cb.queuePtr >= cb.queueSize {
		cb.queuePtr = 0
	}
}

type getUserReq struct{ id int }
type getUserResp struct{ user user }

type createUserReq struct{ user user }
type createUserResp struct{}

type user struct {
}

type client struct {
}

func (c *client) GetUser(ctx context.Context, req getUserReq) (getUserResp, error)
func (c *client) CreateUser(ctx context.Context, req createUserReq) (createUserResp, error)

type execFunc[T, U any] func(context.Context, T) (U, error)

func foo() {
	client1 := &client{}
	client2 := &client{}
	fmt.Println(client1 == client2)
	// var ex execFunc[getUserReq, getUserResp] = client.GetUser
}
