package convert

import (
	"encoding/binary"
)

func U16ToByteBig(n uint16) []byte {
	s := make([]byte, 2)
	binary.BigEndian.PutUint16(s, n)
	return s
}

func U16ToByteLittle(n uint16) []byte {
	s := make([]byte, 2)
	binary.LittleEndian.PutUint16(s, n)
	return s
}
