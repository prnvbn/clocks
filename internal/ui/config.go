package ui

import (
	"encoding/json"
	"fmt"

	"github.com/prnvbn/clocks/internal/tmz"
)

type AppConfig struct {
	DateFmt       string        `yaml:"dateFormat"`
	RowSizes      []int         `yaml:"rowSizes"`
	Timzones      []ClockConfig `yaml:"timezones"`
	CenterEachRow bool          `yaml:"centerEachRow"`
}

type ClockConfig struct {
	Heading string   `yaml:"heading"`
	Zone    tmz.Zone `yaml:"zone"`
	Color   Color    `yaml:"color"`
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
