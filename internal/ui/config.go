package ui

import (
	"encoding/json"
	"fmt"
	"slices"
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/prnvbn/clocks/internal/tmz"
)

type AppConfig struct {
	Layout     LayoutConfig       `yaml:"layout" json:"layout"`
	ClockCfgs  SortedClockConfigs `yaml:"clocks" json:"clocks"`
	Live       bool               `yaml:"live" json:"live"`
	Seconds    bool               `yaml:"seconds" json:"seconds"`
	TwelveHour bool               `yaml:"twelveHour" json:"twelveHour"`
}

type LayoutConfig struct {
	LayoutType    LayoutType `yaml:"layoutType" json:"layoutType"`
	RowSizes      []int      `yaml:"rowSizes" json:"rowSizes"`
	CenterEachRow bool       `yaml:"centerEachRow" json:"centerEachRow"`
}

// NewHorizontalLayout returns a layout with row sizes set to [numClocks]
// i.e. there is only one row which can have numClocks clocks
func NewHorizontalLayout(numClocks int) (l LayoutConfig) {
	l.RowSizes = []int{numClocks}
	l.LayoutType = Horizontal

	return
}

// NewVerticalLayout returns a layout with row sizes set to [1] * numClocks
// i.e. each clock has its own row
func NewVerticalLayout(numClocks int) (l LayoutConfig) {
	l.RowSizes = make([]int, numClocks)
	for i := 0; i < numClocks; i++ {
		l.RowSizes[i] = 1
	}
	l.LayoutType = Vertical

	return
}

// MaxClocks returns the maximum number of clocks
// the current layout supports
func (l LayoutConfig) MaxClocks() (sum int) {
	for _, s := range l.RowSizes {
		sum += s
	}
	return
}

type ClockConfig struct {
	Heading string   `yaml:"heading" json:"heading"`
	Zone    tmz.Zone `yaml:"zone" json:"zone"`
	Color   Color    `yaml:"color" json:"color"`
}

func (c ClockConfig) String() string {
	return c.Color.ToStyle().Sprintf("%s (%s)", c.Heading, c.Zone)
}

// PrettyPrint prints the config in a human readable format
func (c AppConfig) PrettyPrint() error {
	cJSON, err := json.MarshalIndent(c, " ", " ")

	if err != nil {
		return err
	}
	fmt.Println(string(cJSON))

	return nil
}

// SortedClockConfigs is a slice of ClockConfigs
// This slice is always sorted based on the UTC offset of the zones
type SortedClockConfigs []ClockConfig

func (s *SortedClockConfigs) Add(clockCfgs ...ClockConfig) {
	for _, clockCfg := range clockCfgs {
		*s = append(*s, clockCfg)
	}
	// sort by zone
	slices.SortFunc(*s, func(a, b ClockConfig) int {
		return a.Zone.Compare(b.Zone)
	})
}

func (s *SortedClockConfigs) Remove(toRemove ...ClockConfig) {
	*s = slices.DeleteFunc(*s, func(cfg ClockConfig) bool {
		return slices.Contains(toRemove, cfg)
	})
}

func (s SortedClockConfigs) Filter(searchTerm string) (filtered SortedClockConfigs, n int) {
	filtered = make(SortedClockConfigs, 0, len(s))
	for _, clockCfg := range s {
		if fuzzy.Match(strings.ToLower(searchTerm), strings.ToLower(clockCfg.Heading)) {
			filtered = append(filtered, clockCfg)
			n++
		}
	}
	return
}
