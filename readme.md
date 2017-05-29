# Goflake ❄️
A UUID generation package written in Go

> *"No two Gophers are the same"* - Me

That's why we need UUIDs to identify the special snowflakes.

## Documentation
The API docs can be viewed [here](https://godoc.org/github.com/JakeHL/Goflake) or generated with godoc.

### Usage
An example of generating a v4 UUID and outputting it
```go
import (
    "fmt"
    "github.com/jakehl/goflake"
)

func main() {
    v4UUID := goflake.NewV4UUID()
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