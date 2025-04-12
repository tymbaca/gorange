package model

type Packet struct {
	Ver  int8
	Size int16 `bin:"lenof:Data"`
	Data []byte
}
