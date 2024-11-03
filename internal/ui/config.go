package ui

import (
	"encoding/json"
	"fmt"
	"iter"
	"maps"
	"slices"
	"strings"

	"github.com/prnvbn/clocks/internal/match"
	"github.com/prnvbn/clocks/internal/tmz"
)

type AppConfig struct {
	Layout     LayoutConfig `yaml:"layout" json:"layout"`
	ClockCfgs  ClockConfigs `yaml:"clocks" json:"clocks"`
	Live       bool         `yaml:"live" json:"live"`
	Seconds    bool         `yaml:"seconds" json:"seconds"`
	TwelveHour bool         `yaml:"twelveHour" json:"twelveHour"`
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

// ClockConfigs is a set of ClockConfigs
type ClockConfigs map[ClockConfig]struct{}

func (s *ClockConfigs) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// Unmarshal into a temporary slice of strings
	var items []ClockConfig
	if err := unmarshal(&items); err != nil {
		return err
	}

	// Initialize the StringSet map
	*s = make(ClockConfigs)
	for _, item := range items {
		(*s)[item] = struct{}{}
	}

	return nil
}

func (s ClockConfigs) Slice() []ClockConfig {
	return slices.Collect(maps.Keys(s))
}

func (s ClockConfigs) Sorted() iter.Seq2[int, ClockConfig] {
	ccfgs := slices.Collect(maps.Keys(s))
	slices.SortFunc(ccfgs, func(a, b ClockConfig) int {
		return a.Zone.Compare(b.Zone)
	})

	return func(yield func(int, ClockConfig) bool) {
		for i, ccfg := range ccfgs {
			if !yield(i, ccfg) {
				return
			}
		}
	}
}

func (s ClockConfigs) Add(clockCfgs ...ClockConfig) {
	for _, clockCfg := range clockCfgs {
		s[clockCfg] = struct{}{}
	}
}

func (s ClockConfigs) Remove(toRemove ...ClockConfig) {
	for _, toRm := range toRemove {
		delete(s, toRm)
	}
}

func (s ClockConfigs) Filter(searchTerms ...string) (filtered ClockConfigs, n int) {
	filtered = make(ClockConfigs, 0)

	for _, st := range searchTerms {
		for ccfg := range s.fuzzyFiltered(st) {
			filtered[ccfg] = struct{}{}
		}

		n += len(filtered)
	}

	return
}

func (s ClockConfigs) fuzzyFiltered(searchTerm string) iter.Seq[ClockConfig] {
	return func(yield func(ClockConfig) bool) {
		for clockCfg := range s {
			if match.Fuzzy(strings.ToLower(searchTerm), strings.ToLower(clockCfg.Heading)) {
				if !yield(clockCfg) {
					return
				}
			}
		}
	}
}
