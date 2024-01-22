package clocks

type Config struct {
	DateFmt       string `yaml:"dateFormat"`
	RowSizes      []int  `yaml:"rowSizes"`
	Timzones      []TMZ  `yaml:"timezones"`
	CenterEachRow bool   `yaml:"centerEachRow"`
}

type TMZ struct {
	Name     string `yaml:"name"`
	Timezone string `yaml:"timezone"` // enum?
	Color    string `yaml:"color"`    // enum?
	DateFmt  string `yaml:"dateFormat"`
}

// TODO: validate config
func (c Config) Validate() error {
	return nil
}
