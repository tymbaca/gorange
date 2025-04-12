package gateway

import (
	"github.com/anthdm/hollywood/actor"
	"github.com/tymbaca/gorange/internal/game/model"
)

type gateway struct{}

func Send(ctx *actor.Context, pid, player *actor.PID, pack model.Packet)
