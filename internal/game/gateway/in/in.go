package gateway

import (
	"github.com/anthdm/hollywood/actor"
	gateway "github.com/tymbaca/gorange/internal/game/gateway/model"
	"github.com/tymbaca/gorange/internal/game/model"
)

func Send(ctx *actor.Context, pid, player *actor.PID, pack model.Packet) {
	ctx.Send(pid, gateway.OutMsg{To: player, Pack: pack})
}
