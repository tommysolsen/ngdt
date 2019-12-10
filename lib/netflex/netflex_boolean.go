package netflex

import "fmt"

// A Boolean is just a dumb bool as a string
type Boolean bool

func (nf *Boolean) UnmarshalJSON(data []byte) error {
	if len(data) != 3 {
		return fmt.Errorf("Unable to marshal %s into a Netflex bool(size must be 3 bytes)", data)
	}
	x := Boolean(data[1] == 49)
	*nf = x
	return nil
}

func (nf Boolean) MarshalJSON() ([]byte, error) {
	val := "\"0\""
	if nf == true {
		val = "\"1\""
	}
	return []byte(val), nil
}
