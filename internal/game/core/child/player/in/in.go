package player

import (
	"github.com/anthdm/hollywood/actor"
	player "github.com/tymbaca/gorange/internal/game/core/child/player/model"
	"github.com/tymbaca/gorange/internal/game/model"
)

func SendIn(ctx *actor.Context, pid *actor.PID, pack model.Packet) {
	ctx.Send(pid, player.InPlayerMsg{Pack: pack})
}

func SendOut(ctx *actor.Context, pid *actor.PID, pack model.Packet) {
	ctx.Send(pid, player.OutPlayerMsg{Pack: pack})
}
