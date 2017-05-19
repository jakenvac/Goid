package goid

// UUID : Represents a UUID or GUID
type UUID struct {
	Version int
	Octets  [16]byte
}

// ToString : Converts the desired guid into a string
func (g *UUID) ToString() string {
	// TODO Implement UUID to string
	var result string

	var iterations = [5]int{7, 11, 15, 19, 31}

	for _, oc := range g.Octets {
		result += string(oc)
	}

	for _, it := range iterations {
		result = result[:it] + "-" + result[it:]
	}

	return result
}
