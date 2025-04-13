package model

import (
	"log/slog"
	"strings"
)

type Packet struct {
	Ver        int8
	HeaderSize int16 `bin:"lenof:Header"`
	Header     []byte
	BodySize   int16 `bin:"lenof:Body"`
	Body       []byte
}

func ParseHeader(data []byte) map[string]string {
	str := string(data)

	rows := strings.Split(str, "\n")

	header := make(map[string]string, len(rows))
	for _, row := range rows {
		splitted := strings.SplitN(row, ":", 2)
		if len(splitted) != 2 {
			slog.Debug("got incorrect header row", "splitted", splitted)
			continue
		}

		header[splitted[0]] = splitted[1]
	}

	return header
}

type Command = string

const (
	CommandConnect    Command = "connect"
	CommandDisconnect Command = "disconnect"
	CommandMsg        Command = "msg"
)
