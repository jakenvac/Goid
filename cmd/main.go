package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jakehl/goid"
)

func main() {

	countstr := "1"
	if len(os.Args) > 1 {
		countstr = os.Args[1]
	}

	count, err := strconv.Atoi(countstr)
	if err != nil {
		return
	}

	for i := 0; i < count; i++ {
		uuid := goid.NewV4UUID()
		fmt.Println(uuid)
	}
}
