package gateway

import (
	"net"

	"github.com/tymbaca/gorange/internal/game/model"
)

type InMsg struct {
	From string
	Pack model.Packet
}

type OutMsg struct {
	To   string
	Pack model.Packet
}

type ConnectMsg struct {
	ID   string
	Conn net.Conn
}

type DisconnectMsg struct {
	ID string
}
