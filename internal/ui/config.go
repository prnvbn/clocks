package ui

import (
	"encoding/json"
	"fmt"

	"github.com/prnvbn/clocks/internal/tmz"
)

type AppConfig struct {
	DateFmt       string      `yaml:"dateFormat" json:"dateFormat"`
	RowSizes      []int       `yaml:"rowSizes" json:"rowSizes"`
	ClockCfgs     ClockCfgSet `yaml:"timezones" json:"timezones"`
	CenterEachRow bool        `yaml:"centerEachRow" json:"centerEachRow"`
}

type ClockConfig struct {
	Heading string   `yaml:"heading" json:"heading"`
	Zone    tmz.Zone `yaml:"zone" json:"zone"`
	Color   Color    `yaml:"color" json:"color"`
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
