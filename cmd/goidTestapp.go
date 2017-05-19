package main

import (
	"encoding/hex"

	"fmt"

	"github.com/jakehl/goid"
)

func main() {
	bytes := [16]byte{}
	for i := 0; i < 16; i++ {
		byteVals, _ := hex.DecodeString("ff")
		bytes[i] = byteVals[0]
	}

	guid := &goid.UUID{Version: 4, Octets: bytes}
	fmt.Println(guid.ToString())

}
