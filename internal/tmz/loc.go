package tmz

import (
	"fmt"
	"time"
)

type Loc string

func (l Loc) String() string {
	return fmt.Sprintf("%s %s UTC", string(l), l.GetUTCOffset())
}

// GetUTCOffset returns the UTC offset for the given location
func (l Loc) GetUTCOffset() string {
	loc, _ := time.LoadLocation(string(l)) // add err check?
	t := time.Now().In(loc)

	_, offset := t.Zone()

	hours := offset / 3600
	remainder := offset % 3600
	minutes := remainder / 60

	return fmt.Sprintf("%+d:%02d", hours, minutes)
}
