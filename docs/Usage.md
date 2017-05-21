// TODO Flesh this out  
```Go
package main

import (
	"fmt"
	"github.com/jakehl/goid"
)

func main() {
	for i := 0; i < 100; i++ {
		uuid := goid.NewV4UUID()
		fmt.Println(uuid.ToString())
	}
}
```