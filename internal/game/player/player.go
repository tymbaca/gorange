package player

import (
	"github.com/anthdm/hollywood/actor"
	core "github.com/tymbaca/gorange/internal/game/core/in"
	"github.com/tymbaca/gorange/internal/game/gateway"
)

type player struct {
	core    *actor.PID
	gateway *actor.PID
}

func (p *player) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case *inPlayerMsg:
		core.Broadcast(ctx, p.core, msg.pack)
	case *outPlayerMsg:
		gateway.Send(ctx, p.gateway, ctx.PID(), msg.pack)
	}
}
