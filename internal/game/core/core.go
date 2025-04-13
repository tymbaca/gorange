package Core

import (
	"github.com/anthdm/hollywood/actor"
	core "github.com/tymbaca/gorange/internal/game/core/model"
	player "github.com/tymbaca/gorange/internal/game/player/in"
)

type Core struct {
	players []*actor.PID
}

func (c *Core) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case core.BroadcastMsg:
		c.broadcast(ctx, ctx.Sender(), msg)
	}
}

func (c *Core) broadcast(ctx *actor.Context, from *actor.PID, msg core.BroadcastMsg) {
	for _, p := range c.players {
		if p.Equals(from) {
			continue
		}

		player.SendOut(ctx, p, msg.Pack)
	}
}
