package gateway

import (
	"github.com/anthdm/hollywood/actor"
	"github.com/tymbaca/gorange/internal/game/model"
)

type OutMsg struct {
	To   *actor.PID
	Pack model.Packet
}
