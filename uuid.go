package goflake

/* Goflake (c) 2017 jlangford.uk */
/* This file contains the struct and it's relevant methods for a UUID */

import (
	"encoding/hex"
)

// UUID Represents a UUID or GUID
type UUID [16]byte

// ToString converts the UUID into its string representation
func (u *UUID) ToString() string {
	result := hex.EncodeToString(u[:])
	result = result[:8] + "-" + result[8:12] + "-" + result[12:16] + "-" + result[16:20] + "-" + result[20:]
	return result
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

// GetVariant returns the variant of the UUID based on the 17th character (1st char of  9th byte) of the UUID
func (u *UUID) GetVariant() string {
	bytes := []byte{u[8]}
	result := hex.EncodeToString(bytes)[:1]
	return result
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
