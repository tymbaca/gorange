package listener

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"

	"github.com/anthdm/hollywood/actor"
	"github.com/tymbaca/gorange/internal/game/assert"
	gateway "github.com/tymbaca/gorange/internal/game/gateway/in"
	"github.com/tymbaca/gorange/internal/game/model"
	"github.com/tymbaca/sbinary"
)

func New(addr string) actor.Producer {
	return func() actor.Receiver {
		return &Listener{addr: addr}
	}
}

type Listener struct {
	addr string
	done chan struct{}
}

func (l *Listener) Receive(ctx *actor.Context) {
	switch ctx.Message().(type) {
	case actor.Started:
		l.done = make(chan struct{})
		go listen(ctx, l)
	case actor.Stopped:
		close(l.done)
	}
}

func listen(ctx *actor.Context, ll *Listener) {
	l, err := net.Listen("tcp", ll.addr)
	if err != nil {
		panic(err)
	}

	go func() {
		<-ll.done
		err := l.Close()
		if err != nil {
			panic(err)
		}
	}()

	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}

		go func() {
			<-ll.done
			err := conn.Close()
			if err != nil {
				panic(err)
			}
		}()

		go func() {
			err = handleConn(ctx, conn)
			if err != nil {
				panic(err)
			}
		}()
	}
}

func handleConn(ctx *actor.Context, conn net.Conn) error {
	for {
		var pack model.Packet
		err := sbinary.NewDecoder(conn).Decode(&pack, binary.BigEndian)
		if err != nil {
			return fmt.Errorf("decode package: %w", err)
		}

		stop, err := handlePacket(ctx, conn, pack)
		if stop {
			break
		}
	}

	return nil
}

func handlePacket(ctx *actor.Context, conn net.Conn, pack model.Packet) (bool, error) {
	switch pack.Ver {
	case 1:
	default:
		log.Panicf("unsupported version %d", pack.Ver)
	}

	header := model.ParseHeader(pack.Header)

	// TODO: auth
	id, ok := header["id"]
	assert.True(ok, "id header must be set")

	command, ok := header["command"]
	assert.True(ok, "command header must be set")

	switch command {
	case model.CommandConnect:
		gateway.Connect(ctx, ctx.Parent(), id, conn)

	case model.CommandDisconnect:
		gateway.Disconnect(ctx, ctx.Parent(), id)

	case model.CommandMsg:
		gateway.SendIn(ctx, ctx.Parent(), id, pack)

	default:
		assert.Panic("unsupported command %s", command)
	}

	panic("unreachable")
}
