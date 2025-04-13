package gateway

import (
	"encoding/binary"
	"log"
	"log/slog"
	"net"

	"github.com/anthdm/hollywood/actor"
	"github.com/tymbaca/gorange/internal/game/core/child/gateway/child/listener"
	gateway "github.com/tymbaca/gorange/internal/game/core/child/gateway/model"
	player "github.com/tymbaca/gorange/internal/game/core/child/player/in"
	"github.com/tymbaca/sbinary"
)

func New(addr string) actor.Producer {
	return func() actor.Receiver {
		return &Gateway{
			addr:    addr,
			pidMap:  make(map[string]*actor.PID),
			peerMap: make(map[string]net.Conn),
		}
	}
}

type Gateway struct {
	addr    string
	pidMap  map[string]*actor.PID
	peerMap map[string]net.Conn
}

func (g *Gateway) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case actor.Started:
		ctx.SpawnChild(listener.New(g.addr), "listener")

		slog.Info("gateway started", "pid", ctx.PID().String())

	case gateway.InMsg:
		pid, ok := g.pidMap[msg.From]
		if !ok {
			log.Panicf("no pid for id %s", msg.From)
		}

		player.SendIn(ctx, pid, msg.Pack)

	case gateway.OutMsg:
		peer, ok := g.peerMap[msg.To]
		if !ok {
			log.Panicf("no peer for id %s", msg.To)
		}

		err := sbinary.NewEncoder(peer).Encode(msg.Pack, binary.BigEndian)
		if err != nil {
			log.Panicf("can't send packet to peer %s: %s", msg.To, err)
		}

	case gateway.ConnectMsg:
		g.peerMap[msg.ID] = msg.Conn
		// TODO:
		// g.pidMap

	case gateway.DisconnectMsg:
		delete(g.peerMap, msg.ID)
		delete(g.pidMap, msg.ID)
	}
}
