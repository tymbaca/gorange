package player

import (
	"log/slog"

	"github.com/anthdm/hollywood/actor"
	gateway "github.com/tymbaca/gorange/internal/game/core/child/gateway/in"
	player "github.com/tymbaca/gorange/internal/game/core/child/player/model"
	core "github.com/tymbaca/gorange/internal/game/core/in"
)

func New(gateway *actor.PID) actor.Producer {
	return func() actor.Receiver {
		return &Player{
			gateway: gateway,
		}
	}
}

type Player struct {
	gateway *actor.PID
}

func (p *Player) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case actor.Started:
		slog.Info("player started", "pid", ctx.PID().String())

	case player.InPlayerMsg:
		core.BroadcastToPlayers(ctx, ctx.Parent(), msg.Pack)

	case player.OutPlayerMsg:
		gateway.SendOut(ctx, p.gateway, ctx.PID(), msg.Pack)
	}
}
