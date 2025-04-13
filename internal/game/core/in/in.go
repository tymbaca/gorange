package core

import (
	"github.com/anthdm/hollywood/actor"
	core "github.com/tymbaca/gorange/internal/game/core/model"
	"github.com/tymbaca/gorange/internal/game/model"
)

func Broadcast(ctx *actor.Context, pid *actor.PID, pack model.Packet) {
	ctx.Send(pid, core.BroadcastMsg{Pack: pack})
}
