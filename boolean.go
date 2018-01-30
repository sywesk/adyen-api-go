package adyen

import (
	"strconv"
	"strings"
)

// stringBool allows us to unmarhsal Adyen Boolean values
// which appear as strings instead of bools.
type stringBool bool

func (b *stringBool) UnmarshalJSON(data []byte) (err error) {
	str := strings.TrimFunc(strings.ToLower(string(data)), func(c rune) bool {
		return c == ' ' || c == '"'
	})

	parsed, err := strconv.ParseBool(str)
	if err != nil {
		return
	}

	*b = stringBool(parsed)
	return
}