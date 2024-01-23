package clocks

import "github.com/prnvbn/clocks/internal/tmz"

type Config struct {
	DateFmt       string `yaml:"dateFormat"`
	RowSizes      []int  `yaml:"rowSizes"`
	Timzones      []TMZ  `yaml:"timezones"`
	CenterEachRow bool   `yaml:"centerEachRow"`
}

type TMZ struct {
	Name     string   `yaml:"name"`
	Timezone tmz.Zone `yaml:"timezone"`
	Color    Color    `yaml:"color"`
}

// TODO: validate config
func (c Config) Validate() error {
	return nil
}
