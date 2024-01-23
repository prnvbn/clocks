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
