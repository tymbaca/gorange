package core

import (
	"log/slog"

	"github.com/anthdm/hollywood/actor"
	"github.com/tymbaca/gorange/internal/game/core/child/gateway"
	playeractor "github.com/tymbaca/gorange/internal/game/core/child/player"
	player "github.com/tymbaca/gorange/internal/game/core/child/player/in"
	core "github.com/tymbaca/gorange/internal/game/core/model"
)

func New(addr string) actor.Producer {
	return func() actor.Receiver {
		return &Core{
			addr:    addr,
			players: make(map[string]*actor.PID),
		}
	}
}

type Core struct {
	addr    string
	players map[string]*actor.PID
	gateway *actor.PID
}

func (c *Core) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case actor.Started:
		pid := ctx.SpawnChild(gateway.New(c.addr), "gateway")
		c.gateway = pid

		slog.Info("core started", "pid", ctx.PID().String())

	case core.BroadcastMsg:
		c.broadcast(ctx, ctx.Sender(), msg)

	case core.SpawnPlayerMsg:
		pid := ctx.SpawnChild(playeractor.New(c.gateway), msg.ID)
		c.players[msg.ID] = pid
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
