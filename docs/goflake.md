# goflake
--
    import "github.com/jakehl/goflake"

Package goflake provides the tools needed to produce and manipulate UUIDs

## Usage

#### type UUID

```go
type UUID [16]byte
```

UUID Represents a UUID or GUID

#### func  GetUUIDFromString

```go
func GetUUIDFromString(strUUID string) (*UUID, error)
```
GetUUIDFromString takes a string representing a UUID and converts it to a UUID
type Returns an error if the string does not match a UUID

#### func  NewNilUUID

```go
func NewNilUUID() *UUID
```
NewNilUUID returns a nil UUID (All bits set to zero)

#### func  NewV4UUID

```go
func NewV4UUID() *UUID
```
NewV4UUID returns a randomized version 4 UUID

#### func (*UUID) GetVariant

```go
func (u *UUID) GetVariant() string
```
GetVariant returns the variant of the UUID based on the 17th character (1st char
of 9th byte) of the UUID

#### func (*UUID) GetVersion

```go
func (u *UUID) GetVersion() string
```
GetVersion returns the version of the UUID based on the 13th character (1st char
of 7th byte) of the UUID

#### func (*UUID) SetVariant

```go
func (u *UUID) SetVariant()
```
SetVariant sets the variant of the guid (first nibble of the 9th byte) TODO
finetune uuid.SetVariant

#### func (*UUID) SetVersion

```go
func (u *UUID) SetVersion(v byte)
```
SetVersion sets the version of the guid (first nibble of the 7th byte)

#### func (*UUID) ToGUIDString

```go
func (u *UUID) ToGUIDString() string
```
ToGUIDString surrounds the uuid string in { } to mimic a microsoft GUID Note:
This has not been developed to any microsoft standards

#### func (*UUID) ToString

```go
func (u *UUID) ToString() string
```
ToString converts the desired guid into a string
