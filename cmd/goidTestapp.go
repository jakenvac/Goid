package main

import (
	"encoding/hex"

	"fmt"

	"github.com/jakehl/goid"
)

func main() {

	bytes := goid.UUID{}
	for i := 0; i < 16; i++ {
		var byteVals []byte
		var hexVal string
		if i == 6 {
			hexVal = "42"
		} else if i == 8 {
			hexVal = "34"
		} else {
			hexVal = "ff"
		}
		byteVals, _ = hex.DecodeString(hexVal)
		bytes[i] = byteVals[0]
	}

	uuid := &bytes

	fmt.Println("UUID:     " + uuid.ToString())
	fmt.Println("Version:  " + uuid.GetVersion())
	uuid.SetVersion(3)
	fmt.Println("Reset Version:  " + uuid.GetVersion())
	fmt.Println("Variant:  " + uuid.GetVariant() + "\n")

	nilUUID := goid.MakeNilUUID()
	fmt.Println("Nil UUID: " + nilUUID.ToString() + "\n")

	///////////////////////////////////
	// byte1 := byte(125)
	// byte1hex := hex.EncodeToString([]byte{byte1})
	// fmt.Println(byte1hex)

	// byte2 := byte(0x40)
	// byte2hex := hex.EncodeToString([]byte{byte2})
	// fmt.Println(byte2hex)

	// byte3 := byte(0xf)
	// byte3hex := hex.EncodeToString([]byte{byte3})
	// fmt.Println(byte3hex + "\n")

	// byte1 = (byte1 & byte3)
	// byte1 = (byte1 | 0x40)

	// byteResult := hex.EncodeToString([]byte{byte1})
	// fmt.Println(byteResult)

	///////////////////////////////////

}
