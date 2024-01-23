package ui

import (
	"math"
	"os"
	"slices"

	"github.com/prnvbn/clocks/internal/tmz"
	"github.com/pterm/pterm"
	"golang.org/x/exp/maps"
	"golang.org/x/term"
)

var (
	colorToStyle = map[string]*pterm.Style{
		Black.String():   pterm.FgBlack.ToStyle(),
		Red.String():     pterm.FgRed.ToStyle(),
		Green.String():   pterm.FgGreen.ToStyle(),
		Yellow.String():  pterm.FgYellow.ToStyle(),
		Blue.String():    pterm.FgBlue.ToStyle(),
		Magenta.String(): pterm.FgMagenta.ToStyle(),
		Cyan.String():    pterm.FgCyan.ToStyle(),
		White.String():   pterm.FgWhite.ToStyle(),
	}
)

const (
	minMaxHeight = 5
)

// TODO? breakdown further?
func SelectClocks() []ClockConfig {
	cntries := maps.Keys(tmz.CountryZonesMap)
	slices.Sort(cntries)

	cntryMenuHeading := pterm.ThemeDefault.PrimaryStyle.Sprint("Please select countries; you can select the timezones after this ") + pterm.ThemeDefault.SecondaryStyle.Sprint("[type to search]")
	selectedCntries, _ := pterm.DefaultInteractiveMultiselect.
		WithOptions(cntries).
		WithDefaultText(cntryMenuHeading).
		WithClearAllEnabled(true).
		WithSelectAllEnabled(false).
		WithMaxHeight(maxHeight()).
		Show()

	var zones []tmz.Zone
	for _, cntry := range selectedCntries {
		cntryToZones := tmz.CountryZonesMap[cntry]
		zones = append(zones, cntryToZones...)
	}

	selectedZones := SelectZones(zones)
	var clockCfgs []ClockConfig

	for _, z := range selectedZones {
		// TODO? show sample number in an area next to the select menu
		color, _ := pterm.NewGenericInteractiveSelect[Color]().
			WithMaxHeight(maxHeight()).
			WithOptions(Colors).
			WithDefaultText("Select a color for " + pterm.Bold.Sprint(z)).
			WithRenderSelectedOptionFunc(func(s string) string {
				return colorToStyle[s].
					Add(*pterm.Bold.ToStyle()).
					Sprintf("  %s\n", s)
			}).
			Show()

		textInput := pterm.DefaultInteractiveTextInput.
			WithMultiLine(false).
			WithDefaultText("Enter the display name for " + pterm.Bold.Sprint(z)).
			WithDefaultValue(z.City())

		heading, _ := textInput.Show()

		clockCfgs = append(clockCfgs, ClockConfig{
			Heading: heading,
			Zone:    z,
			Color:   color,
		})
	}

	return clockCfgs

}

func maxHeight() int {
	_, h, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		h = minMaxHeight
	}
	return int(math.Max(1, float64(h-5)))
}