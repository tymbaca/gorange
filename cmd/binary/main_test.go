package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncodeDecode(t *testing.T) {
	req := Request{
		MessageSize: 1,
		Header: Header{
			Version:       3,
			CorrelationID: 2,
			// Body: []byte("hello"),
			// Msg: "world",
		},
	}

	buf := bytes.NewBuffer(nil)
	require.Nil(t, NewEncoder(buf).Encode(req, binary.BigEndian))

	fmt.Println(buf.Bytes())

	var reqDecoded Request
	require.NotNil(t, NewDecoder(buf).Decode(reqDecoded, binary.BigEndian)) // not pointer
	require.Nil(t, NewDecoder(buf).Decode(&reqDecoded, binary.BigEndian))

	require.Equal(t, req, reqDecoded)
}

type Request struct {
	MessageSize uint32 `bin:"lenofrest"`
	Header      Header
}

type Header struct {
	Version       byte
	CorrelationID int32
	BodySize      uint32 `bin:"lenof:Body"`
	// Body          []byte
	// Msg           string
}
