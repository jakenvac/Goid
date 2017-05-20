package main

import (
	"encoding/hex"

	"fmt"

	"github.com/jakehl/goid"
)

func main() {

	bytes := [16]byte{}
	for i := 0; i < 16; i++ {
		var byteVals []byte
		var hexVal string
		if i == 6 {
			hexVal = "12"
		} else if i == 8 {
			hexVal = "34"
		} else {
			hexVal = "ff"
		}
		byteVals, _ = hex.DecodeString(hexVal)
		bytes[i] = byteVals[0]
	}

	uuid := &goid.UUID{Octets: bytes}

	fmt.Println("UUID:     " + uuid.ToString())
	fmt.Println("Version:  " + uuid.GetVersion())
	fmt.Println("Variant:  " + uuid.GetVariant())

	nilUUID := goid.MakeNilUUID()
	fmt.Println("Nil UUID: " + nilUUID.ToString())

}
