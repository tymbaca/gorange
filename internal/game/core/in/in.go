package core

import (
	"github.com/anthdm/hollywood/actor"
	"github.com/tymbaca/gorange/internal/game/model"
)

func Broadcast(ctx *actor.Context, pid *actor.PID, pack model.Packet) {
	ctx.Send(pid, pack)
}
