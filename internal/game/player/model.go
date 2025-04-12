package player

import "github.com/tymbaca/gorange/internal/game/model"

type inPlayerMsg struct {
	pack model.Packet
}

type outPlayerMsg struct {
	pack model.Packet
}
