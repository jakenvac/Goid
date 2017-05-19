package main

import (
	"fmt"

	"github.com/jakehl/goid"
)

func main() {
	a := &goid.UUID{}
	fmt.Println(a)
}
