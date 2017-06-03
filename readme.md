# Goid
A UUID generation package written in Go  

[![Go Report Card](https://goreportcard.com/badge/github.com/JakeHL/goid)](https://goreportcard.com/report/github.com/JakeHL/goid)
[![cover.run go](https://cover.run/go/github.com/jakehl/goid.svg)](https://cover.run/go/github.com/jakehl/goid)
[![Go Docs](https://godoc.org/github.com/jakehl/goid?status.svg)](https://godoc.org/github.com/jakehl/goid)

## Documentation
The API docs can be viewed [here](https://godoc.org/github.com/JakeHL/goid) or generated with godoc.

### Usage
An example of generating a v4 UUID and outputting it
```go
import (
    "fmt"
    "github.com/jakehl/goid"
)

func main() {
    v4UUID := goid.NewV4UUID()
    fmt.PrintLn(v4UUID)
}
```

## Notes
### Contributing & Package status
Currently this package should generate RFC4122 compliant Version 4 UUIDs and the existing APIs to do so should not change. New features may be added via pull requests, however they should not touch existing APIs

### Todo
- [ ] Add optimised bulk UUID generation
- [ ] Add UUID math operations
- [ ] Add other UUID versions? 

### References
The following sources were referenced during the development of this project
- http://www.cryptosys.net/pki/uuid-rfc4122.html
- http://www.ietf.org/rfc/rfc4122.txt
- https://en.wikipedia.org/wiki/Universally_unique_identifier