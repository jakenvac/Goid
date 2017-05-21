package main

import (
	"fmt"

	"github.com/jakehl/goid"
)

func main() {
	for i := 0; i < 100; i++ {
		uuid := goid.NewV5UUID()
		fmt.Println(uuid.ToString())
	}
}
