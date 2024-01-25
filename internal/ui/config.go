package ui

import (
	"encoding/json"
	"fmt"
	"slices"

	"github.com/prnvbn/clocks/internal/tmz"
)

type AppConfig struct {
	DateFmt   string             `yaml:"dateFormat" json:"dateFormat"`
	Layout    LayoutConfig       `yaml:"layout" json:"layout"`
	ClockCfgs SortedClockConfigs `yaml:"timezones" json:"timezones"`
	Live      bool               `yaml:"live" json:"live"`
}

type LayoutConfig struct {
	RowSizes      []int `yaml:"rowSizes" json:"rowSizes"`
	CenterEachRow bool  `yaml:"centerEachRow" json:"centerEachRow"`
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
	return fmt.Sprintf("%s (%s)", c.Heading, c.Zone)
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
