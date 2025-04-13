package player

import (
	"github.com/anthdm/hollywood/actor"
	core "github.com/tymbaca/gorange/internal/game/core/in"
	gateway "github.com/tymbaca/gorange/internal/game/gateway/in"
	player "github.com/tymbaca/gorange/internal/game/player/model"
)

func New()

type Player struct {
	core    *actor.PID
	gateway *actor.PID
}

func (p *Player) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case player.InPlayerMsg:
		core.Broadcast(ctx, p.core, msg.Pack)
	case player.OutPlayerMsg:
		gateway.SendIn(ctx, p.gateway, ctx.PID(), msg.Pack)
	}
}
