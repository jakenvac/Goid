package main

import (
	"fmt"

	"github.com/jakehl/goflake"
)

func main() {
	for i := 0; i < 1000000; i++ {
		uuid := goflake.NewV4UUID()
		fmt.Println(uuid.ToString())
	}
}
