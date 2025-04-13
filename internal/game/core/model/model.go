package core

import (
	"github.com/tymbaca/gorange/internal/game/model"
)

type BroadcastMsg struct {
	Pack model.Packet
}

type SpawnPlayerMsg struct {
	ID string
}
