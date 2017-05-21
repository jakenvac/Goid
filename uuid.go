package goid

/* Goid (c) 2017 jlangford.uk */
/* This file contains the struct and it's relevant methods for a UUID */

import (
	"encoding/hex"
	"errors"
)

// UUID Represents a UUID or GUID
type UUID [16]byte

// ToString converts the desired guid into a string
func (u *UUID) ToString() string {
	result := hex.EncodeToString(u[:])
	result = result[:8] + "-" + result[8:12] + "-" + result[12:16] + "-" + result[16:20] + "-" + result[20:]
	return result
}

// ToGUIDString surrounds the uuid string in { } to mimic a microsoft guid
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
// Returns an error if the version is not valid
func (u *UUID) SetVersion(v uint16) error {
	if v > 0 && v <= 5 {
		version := byte(v) << 4
		u[6] = version | (u[6] & 0xf)
		return nil
	}
	return errors.New("UUID version outside of range")
}

// SetVariant sets the variant of the guid (first nibble of the 9th byte)
// Returns an error if the variant not valid
// TODO finetune uuid.SetVariant
func (u *UUID) SetVariant(v uint16) error {
	if v > 0 && v <= 5 {
		version := byte(v) << 4
		u[6] = version | (u[8] & 0xf)
		return nil
	}
	return errors.New("UUID version outside of range")
}
