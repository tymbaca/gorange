package player

import (
	"github.com/anthdm/hollywood/actor"
	"github.com/tymbaca/gorange/internal/game/model"
	player "github.com/tymbaca/gorange/internal/game/player/model"
)

func SendIn(ctx *actor.Context, pid *actor.PID, pack model.Packet) {
	ctx.Send(pid, player.InPlayerMsg{Pack: pack})
}

func SendOut(ctx *actor.Context, pid *actor.PID, pack model.Packet) {
	ctx.Send(pid, player.OutPlayerMsg{Pack: pack})
}
