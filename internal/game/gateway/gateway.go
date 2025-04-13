package gateway

import (
	"encoding/binary"
	"io"
	"log"
	"sync"

	"github.com/anthdm/hollywood/actor"
	gateway "github.com/tymbaca/gorange/internal/game/gateway/model"
	"github.com/tymbaca/sbinary"
)

type Gateway struct {
	mu      sync.Mutex
	peerMap map[string]io.Writer
}

func (g *Gateway) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case gateway.OutMsg:
		peer, ok := g.peerMap[msg.To.ID]
		if !ok {
			log.Panicf("no peer for id %s", msg.To.ID)
		}

		err := sbinary.NewEncoder(peer).Encode(msg.Pack, binary.BigEndian)
		if err != nil {
			log.Panicf("can't send packet to peer %s: %s", msg.To.ID, err)
		}
	}
}
