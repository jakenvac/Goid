// Package goid provides the tools needed to produce and manipulate V4 UUIDs
package goid

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
)

func init() {
	v4Regex, _ = regexp.Compile(V4UUIDRegex)
}

// V4UUIDRegex is a simple string regex that will match valid v4 UUIDs
const V4UUIDRegex = "\\b[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}\\b"

var v4Regex *regexp.Regexp

// UUID Represents a UUID
type UUID [16]byte

// ToString converts the UUID into its string representation
func (u *UUID) String() string {
	result := hex.EncodeToString(u[:])
	result = result[:8] + "-" + result[8:12] + "-" + result[12:16] + "-" + result[16:20] + "-" + result[20:]
	return result
}

// GetVersion returns the version of the UUID based on the 13th character (1st char of 7th byte) of the UUID
func (u *UUID) GetVersion() string {
	bytes := []byte{u[6]}
	result := hex.EncodeToString(bytes)[:1]
	return result
}

// GetUUIDFromString takes a string representing a UUID and converts it to a UUID type
// Returns an error if the string does not match a UUID
func GetUUIDFromString(strUUID string) (*UUID, error) {
	if ismatch := v4Regex.MatchString(strUUID); !ismatch {
		return nil, fmt.Errorf("%s is not a valid v4 UUID", strUUID)
	}
	strUUID = strings.Replace(strUUID, "-", "", 4)
	resultUUIDSlice, _ := hex.DecodeString(strUUID)
	resultUUID, _ := GetUUIDFromByteSlice(resultUUIDSlice)
	return resultUUID, nil
}

// GetUUIDFromByteSlice takes a slice of bytes and converts it to a UUID
// If the slice is not 16 bytes long it will return an error
// @TODO add version checking
func GetUUIDFromByteSlice(slice []byte) (*UUID, error) {
	if length := len(slice); length != 16 {
		return nil, fmt.Errorf("%s is %d bytes. Expects 16", slice, length)
	}
	uuid := UUID{}
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
func NewV4UUID() (result *UUID) {
	bytes := make([]byte, 16)
	rand.Read(bytes)

	result = &UUID{}
	for i, v := range bytes {
		result[i] = v
	}

	result.setVersion(4)
	result.setVariant()
	return
}

func (u *UUID) Equals(comp *UUID) bool {
	for i, v := range u {
		if comp[i] != v {
			return false
		}
	}
	return true
}

// SetVersion sets the version of the uuid (first nibble of the 7th byte)
func (u *UUID) setVersion(v byte) {
	u[6] = (v << 4) | (u[6] & 0xf)
}

// setVariant sets the variant of the uuid (first nibble of the 9th byte)
func (u *UUID) setVariant() {
	// Clear the first two bits of the byte
	u[8] = u[8] & 0x3f
	// Set the first two bits of the byte to 10
	u[8] = u[8] | 0x80
}
