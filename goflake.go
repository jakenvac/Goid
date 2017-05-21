// Package goflake provides the tools needed to produce and manipulate UUIDs
package goflake

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// GetUUIDFromString takes a string representing a UUID and converts it to a UUID type
// Returns an error if the string does not match a UUID
// BUG(JakeHL) GetUUIDFromString Not Implemented yet
func GetUUIDFromString(strUUID string) (*UUID, error) {
	var resError error
	resUUID := &UUID{}

	// TODO Do stuff in here
	// TODO If error:
	resError = fmt.Errorf("String: %v does not match thr format of a UUID", strUUID)
	return resUUID, resError
}

// NewNilUUID returns a nil UUID (All bits set to zero)
func NewNilUUID() *UUID {
	var bytes UUID = [16]byte{}
	for i := 0; i < 16; i++ {
		stringByte, _ := hex.DecodeString("00")
		bytes[i] = stringByte[0]
	}
	return &bytes
}

// NewV1UUID returns a time and node ID based version 1 UUID
// Note: Upon running, it will check if a NodeID has been initialized, if not, it will do do so.
// This can be done manually with InitializeNewNodeID
func NewV1UUID() *UUID {

}

// NewV4UUID returns a randomized version 4 UUID
func NewV4UUID() *UUID {
	bytes := make([]byte, 16)
	rand.Read(bytes)

	result := UUID{}
	for i, v := range bytes {
		result[i] = v
	}

	result.SetVersion(4)
	result.SetVariant()
	return &result
}

/* == Unexported methods & variables == */
var nodeID [6]byte

// setMulticaseBitOfNodeID sets the last bit of the first octet of the node ID to 1
func setMulticastBitOfNodeID(bytes *[6]byte) {
	bytes[0] = bytes[0] | 0x1
}
