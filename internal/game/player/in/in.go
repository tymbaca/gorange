package player

import (
	"github.com/anthdm/hollywood/actor"
	"github.com/tymbaca/gorange/internal/game/model"
)

func SendIn(ctx *actor.Context, pid *actor.PID, pack model.Packet) {
	ctx.Send(pid, inPlayerMsg{pack: pack})
}

func SendOut(ctx *actor.Context, pid *actor.PID, pack model.Packet) {
	ctx.Send(pid, outPlayerMsg{pack: pack})
}
