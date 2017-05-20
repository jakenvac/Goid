package goid

/* Goid (c) 2017 jlangford.uk */
/* This file contains the struct and it's relevant methods for a UUID */

import (
	"encoding/hex"
)

// UUID Represents a UUID or GUID
type UUID struct {
	Octets [16]byte
}

// ToString converts the desired guid into a string
func (u *UUID) ToString() string {
	result := hex.EncodeToString(u.Octets[:])
	result = result[:8] + "-" + result[8:12] + "-" + result[12:16] + "-" + result[16:20] + "-" + result[20:]
	return result
}

// GetVersion returns the version of the UUID based on the 13th character (1st char of 7th byte) of the UUID
func (u *UUID) GetVersion() string {
	bytes := []byte{u.Octets[6]}
	result := hex.EncodeToString(bytes)[:1]
	return result
}

// GetVariant returns the variant of the UUID based on the 17th character (1st char of  9th byte) of the UUID
func (u *UUID) GetVariant() string {
	bytes := []byte{u.Octets[8]}
	result := hex.EncodeToString(bytes)[:1]
	return result
}
