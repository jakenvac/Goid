package goid

import (
	"encoding/hex"
)

// MakeNilUUID returns a UUID with all bits set to zero
func MakeNilUUID() *UUID {
	bytes := [16]byte{}
	for i := 0; i < 16; i++ {
		stringByte, _ := hex.DecodeString("00")
		bytes[i] = stringByte[0]
	}
	return &UUID{bytes}
}
