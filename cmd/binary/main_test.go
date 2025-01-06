package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncodeDecode(t *testing.T) {
	h := Header{
		Version:       3,
		MessageSize:   1,
		CorrelationID: 2,
		// Body: []byte("hello"),
		// Msg: "world",
	}

	buf := bytes.NewBuffer(nil)

	require.Nil(t, NewEncoder(buf).Encode(h, binary.BigEndian))

	fmt.Println(buf.Bytes())

	var hDecode Header
	require.NotNil(t, NewDecoder(buf).Decode(hDecode, binary.BigEndian)) // not pointer

	require.Nil(t, NewDecoder(buf).Decode(&hDecode, binary.BigEndian))

	require.Equal(t, h, hDecode)
}

type Header struct {
	Version       byte
	MessageSize   uint32 `bin:"lenofall"`
	CorrelationID int32
	BodySize      uint32 `bin:"lenof:Body"`
	Body          []byte
	Msg           string
}
