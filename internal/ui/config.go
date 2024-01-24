package ui

import (
	"encoding/json"
	"fmt"

	"github.com/prnvbn/clocks/internal/tmz"
)

type AppConfig struct {
	DateFmt   string       `yaml:"dateFormat" json:"dateFormat"`
	Layout    LayoutConfig `yaml:"layout" json:"layout"`
	ClockCfgs ClockCfgSet  `yaml:"timezones" json:"timezones"`
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

// TODO: validate config
func (c AppConfig) Validate() error {
	return nil
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

type ClockCfgSet map[ClockConfig]struct{}

func (s *ClockCfgSet) Add(clockCfgs ...ClockConfig) {
	for _, clockCfg := range clockCfgs {
		(*s)[clockCfg] = struct{}{}
	}
}

func (s *ClockCfgSet) Remove(clockCfgs ...ClockConfig) {
	for _, clockCfg := range clockCfgs {
		delete(*s, clockCfg)
	}
}

// UnmarshalYAML implements the yaml.Unmarshaler interface
func (s *ClockCfgSet) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var clocks []ClockConfig
	err := unmarshal(&clocks)
	if err != nil {
		return err
	}

	*s = make(ClockCfgSet)
	for _, clock := range clocks {
		(*s)[clock] = struct{}{}
	}

	return nil
}

// MarshalYAML implements the yaml.Marshaler interface
func (s ClockCfgSet) MarshalYAML() (interface{}, error) {
	var clocks []ClockConfig
	for clock := range s {
		clocks = append(clocks, clock)
	}

	return clocks, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (s *ClockCfgSet) UnmarshalJSON(b []byte) error {
	var clocks []ClockConfig
	err := json.Unmarshal(b, &clocks)
	if err != nil {
		return err
	}

	*s = make(ClockCfgSet)
	for _, clock := range clocks {
		(*s)[clock] = struct{}{}
	}

	return nil
}

// MarshalJSON implements the json.Marshaler interface
func (s ClockCfgSet) MarshalJSON() ([]byte, error) {
	var clocks []ClockConfig
	for clock := range s {
		clocks = append(clocks, clock)
	}

	return json.Marshal(clocks)
}
