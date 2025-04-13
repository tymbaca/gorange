package player

import "github.com/tymbaca/gorange/internal/game/model"

type InPlayerMsg struct {
	Pack model.Packet
}

type OutPlayerMsg struct {
	Pack model.Packet
}
