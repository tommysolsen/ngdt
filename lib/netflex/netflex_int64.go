package netflex

import (
	"fmt"
	"strconv"
)

// NetflexInt64 exists because Netflex has stringified ints
type Int64 int64

// UnmarshalJSON takes a stringified int and makes it a real boy
func (nf *Int64) UnmarshalJSON(data []byte) error {
	if data[0] != 34 || data[len(data)-1] != 34 {
		return fmt.Errorf("Int64 requires marshalled data to be a string, got %s", data)
	}
	if (len(data)) <= 2 {
		*nf = 2
		return nil
	}

	i, err := strconv.ParseInt(string(data[1:len(data)-1]), 10, 64)
	if err != nil {
		*nf = 0
		return nil
	}
	x := Int64(i)
	*nf = x
	return err
}

// MarshalJSON marshals a Int64 int into an stringified Int
func (nf Int64) MarshalJSON() ([]byte, error) {
	return []byte("\"" + strconv.FormatInt(int64(nf), 10) + "\""), nil
}
