// Package goid provides the tools needed to produce and manipulate V4 UUIDs
package goid

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
)

// UUID Represents a UUID or GUID
type UUID [16]byte

// ToString converts the UUID into its string representation
func (u *UUID) ToString() string {
	result := hex.EncodeToString(u[:])
	result = result[:8] + "-" + result[8:12] + "-" + result[12:16] + "-" + result[16:20] + "-" + result[20:]
	return result
}

// String provides stringer functionality for fmt.Print functionality
func (u *UUID) String() string {
	return u.ToString()
}

// ToGUIDString surrounds the UUID string in { } to mimic a microsoft GUID
// Note: This has not been developed to any microsoft standards
func (u *UUID) ToGUIDString() string {
	return "{" + u.ToString() + "}"
}

// GetVersion returns the version of the UUID based on the 13th character (1st char of 7th byte) of the UUID
func (u *UUID) GetVersion() string {
	bytes := []byte{u[6]}
	result := hex.EncodeToString(bytes)[:1]
	return result
}

/* ### UUID Generation ### */

// GetUUIDFromString takes a string representing a UUID and converts it to a UUID type
// Returns an error if the string does not match a UUID
func GetUUIDFromString(strUUID string) (*UUID, error) {
	if ismatch, _ := regexp.MatchString(V4UUIDRegex, strUUID); !ismatch {
		return nil, fmt.Errorf("%s is not a valid v4 UUID", strUUID)
	}
	strUUID = strings.Replace(strUUID, "-", "", 4)
	resultUUIDSlice, _ := hex.DecodeString(strUUID)
	resultUUID, _ := GetUUIDFromByteSlice(resultUUIDSlice)
	return resultUUID, nil
}

// GetUUIDFromByteSlice takes a slice of bytes and converts it to a UUID
// If the slice is not 16 bytes long it will return an error
// This does not check for Version or Variant compliance - Use GetUUIDFromString where possible
func GetUUIDFromByteSlice(slice []byte) (*UUID, error) {
	if len(slice) != 16 {
		return nil, fmt.Errorf("%s is not 16 bytes in length", slice)
	}
	var uuid UUID
	for i, v := range slice {
		uuid[i] = v
	}
	return &uuid, nil
}

// NewNilUUID returns a nil UUID (All bits set to zero)
func NewNilUUID() *UUID {
	return &UUID{}
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

// SetVersion sets the version of the guid (first nibble of the 7th byte)
func (u *UUID) SetVersion(v byte) {
	u[6] = (v << 4) | (u[6] & 0xf)
}

// SetVariant sets the variant of the guid (first nibble of the 9th byte)
// TODO finetune uuid.SetVariant
func (u *UUID) SetVariant() {
	// Clear the first two bits of the byte
	u[8] = u[8] & 0x3f
	// Set the first two bits of the byte to 10
	u[8] = u[8] | 0x80
}

// V4UUIDRegex is a simple string regex that will match valid v4 UUIDs
const V4UUIDRegex = "\\b[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}\\b"
