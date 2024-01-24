package tmz

import (
	"fmt"
	"strings"
	"time"
)

type Zone string

// String returns the location and UTC offset
func (l Zone) String() string {
	return fmt.Sprintf("%s %s UTC", string(l), l.GetUTCOffset())
}

// City returns the city for the given location
// e.g. "America/New_York" -> "New York",  "America/Argentina/Catamarca" -> "Catamarca"
func (l Zone) City() string {
	parts := strings.Split(string(l), "/")
	return strings.Replace(parts[len(parts)-1], "_", " ", -1)
}

// GetUTCOffset returns the UTC offset for the given location
func (l Zone) GetUTCOffset() string {
	loc, _ := time.LoadLocation(string(l)) // add err check?
	t := time.Now().In(loc)

	_, offset := t.Zone()

	hours := offset / 3600
	remainder := offset % 3600
	minutes := remainder / 60

	return fmt.Sprintf("%+d:%02d", hours, minutes)
}

// Compare returns an integer comparing two Zones based on their UTC offset.
// The result will be
// 1. 0 if z == other,
// 2. -1 if z < other, and
// 3. +1 if z > other.
func (z Zone) Compare(other Zone) int {
	thisOffset := z.GetUTCOffset()
	otherOffset := other.GetUTCOffset()

	// "+" > "-" is false in Go
	if thisOffset[0] > otherOffset[0] {
		return -1
	} else if thisOffset[0] < otherOffset[0] {
		return 1
	}

	return strings.Compare(thisOffset, otherOffset)
}

// marshalJSON
func (z Zone) MarshalJSON() ([]byte, error) {
	return []byte(`"` + string(z) + `"`), nil
}

func (z *Zone) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	err := unmarshal(&s)
	if err != nil {
		return err
	}

	*z = Zone(s)
	return nil
}
