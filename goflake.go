// Package goflake provides the tools needed to produce and manipulate UUIDs
package goflake

import (
	"crypto/rand"
)

// InitializeNewRandomNodeID initializes the node ID to a fake MAC address used in
// UUID V1 & V2 generation
func InitializeNewRandomNodeID() {
	bytes := make([]byte, 6)
	rand.Read(bytes)
	for i, v := range bytes {
		nodeID[i] = v
	}
	setMulticastAndLocalBitsOfNodeID()
	isNodeIDSet = true
}

// SetNodeID allows the node ID to be set manually, eg. to a real hardware address
// if overwriteBits is set to true, the last two bits of the first octet will be set to 1
func SetNodeID(ID [6]byte, overwriteBits bool) {
	nodeID = ID
	if overwriteBits {
		setMulticastAndLocalBitsOfNodeID()
	}
	isNodeIDSet = true
}

// GetUUIDFromString takes a string representing a UUID and converts it to a UUID type
// Returns an error if the string does not match a UUID
func GetUUIDFromString(strUUID string) (*UUID, error) {
	// BUG(JakeHL) GetUUIDFromString Not Implemented yet
	panic("Not Implemented")
}

// NewNilUUID returns a nil UUID (All bits set to zero)
func NewNilUUID() *UUID {
	return &UUID{}
}

// NewV1UUID returns a time and node ID based version 1 UUID
// Note: Upon running, it will check if a NodeID has been initialized, if not, it will do do so.
// This can be done manually with InitializeNewRandomNodeID or SetNodeID
func NewV1UUID() *UUID {
	if !isNodeIDSet {
		InitializeNewRandomNodeID()
	}
	// TODO Implement NewV1UUID()
	return nil
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
var isNodeIDSet = false
var nodeID [6]byte

// setMulticaseBitOfNodeID sets the last two bits of the first octet of the node ID to 1
func setMulticastAndLocalBitsOfNodeID() {
	nodeID[0] = nodeID[0] | 0x3
}
