package gateway

import (
	"encoding/binary"
	"log"
	"net"

	"github.com/anthdm/hollywood/actor"
	"github.com/tymbaca/gorange/internal/game/gateway/child/listener"
	gateway "github.com/tymbaca/gorange/internal/game/gateway/model"
	player "github.com/tymbaca/gorange/internal/game/player/in"
	"github.com/tymbaca/sbinary"
)

type Gateway struct {
	core *actor.PID

	pidMap  map[string]*actor.PID
	peerMap map[string]net.Conn
}

func (g *Gateway) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case actor.Started:
		ctx.SpawnChild(listener.New(":8080"), "listener")

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
		g.pidMap

	case gateway.DisconnectMsg:
		delete(g.peerMap, msg.ID)
		delete(g.pidMap, msg.ID)
	}
}
