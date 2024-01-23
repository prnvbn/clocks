package ui

import "github.com/pterm/pterm"

type Color int

const (
	Black Color = iota
	Blue
	Cyan
	Green
	Magenta
	Red
	White
	Yellow
)

var (
	Colors        = []Color{Black, Blue, Cyan, Green, Magenta, Red, White, Yellow}
	colors        = []string{"black", "blue", "cyan", "green", "magenta", "red", "white", "yellow"}
	stringToColor = map[string]Color{
		"black":   Black,
		"red":     Red,
		"green":   Green,
		"yellow":  Yellow,
		"blue":    Blue,
		"magenta": Magenta,
		"cyan":    Cyan,
		"white":   White,
	}
	colorStyles = []*pterm.Style{
		pterm.FgBlack.ToStyle(),
		pterm.FgBlue.ToStyle(),
		pterm.FgCyan.ToStyle(),
		pterm.FgGreen.ToStyle(),
		pterm.FgMagenta.ToStyle(),
		pterm.FgRed.ToStyle(),
		pterm.FgWhite.ToStyle(),
		pterm.FgYellow.ToStyle(),
	}
)

func (c Color) String() string {
	return colors[c]
}
func ColorFromString(s string) Color {
	return stringToColor[s]
}

func (c Color) ToStyle() *pterm.Style {
	return colorStyles[c]
}

// MarshalJSON
func (c Color) MarshalJSON() ([]byte, error) {
	return []byte(`"` + c.String() + `"`), nil
}

func (c Color) MarshalYAML() (interface{}, error) {
	return c.String(), nil
}

// unmarschalYAML
func (c *Color) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	err := unmarshal(&s)
	if err != nil {
		return err
	}

	*c = ColorFromString(s)
	return nil
}
