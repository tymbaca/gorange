package core

import (
	"github.com/anthdm/hollywood/actor"
	"github.com/tymbaca/gorange/internal/game/player"
)

type core struct {
	players []*actor.PID
}

func (c *core) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case *broadcastMsg:
		c.broadcast(ctx, ctx.Sender(), msg)
	}
}

func (c *core) broadcast(ctx *actor.Context, from *actor.PID, msg *broadcastMsg) {
	for _, p := range c.players {
		if p.Equals(from) {
			continue
		}

		player.SendOut(ctx, p, msg.pack)
	}
}
