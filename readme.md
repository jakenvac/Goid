# Goflake ❄️
A UUID generation package written in Go

> *"No two Gophers are the same"* - Me

That's why we need UUIDs to identify the special snowflakes.

## Documentation
Currently docs can be viewed [here](https://godoc.org/github.com/JakeHL/Goflake) or generated with godoc.

### Usage
An example of generating a v4 UUID
```go
import (
    "fmt"
    "github.com/jakehl/goflake"
)

func main() {
    v4UUID := goflake.NewV4UUID()
    fmt.PrintLn(v4UUID.ToString())
}
```

## Notes
This package is under heavy development and is currently for my own learning. If you have ended up here, please use this with caution as
i'll most definitely implement breaking changes during the early phases of development & UUIDs may not be to standard. 🙃

### Sources
The following sources were referenced during the development of this project
- http://www.cryptosys.net/pki/uuid-rfc4122.html
- http://www.ietf.org/rfc/rfc4122.txt
- https://en.wikipedia.org/wiki/Universally_unique_identifier