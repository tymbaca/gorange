package gateway

import (
	"net"

	"github.com/anthdm/hollywood/actor"
	gateway "github.com/tymbaca/gorange/internal/game/core/child/gateway/model"
	"github.com/tymbaca/gorange/internal/game/model"
)

func SendIn(ctx *actor.Context, pid *actor.PID, from string, pack model.Packet) {
	ctx.Send(pid, gateway.InMsg{From: from, Pack: pack})
}

func SendOut(ctx *actor.Context, pid, player *actor.PID, pack model.Packet) {
	ctx.Send(pid, gateway.OutMsg{To: player.ID, Pack: pack})
}

func Connect(ctx *actor.Context, pid *actor.PID, id string, conn net.Conn) {
	ctx.Send(pid, gateway.ConnectMsg{ID: id, Conn: conn})
}

func Disconnect(ctx *actor.Context, pid *actor.PID, id string) {
	ctx.Send(pid, gateway.DisconnectMsg{ID: id})
}
