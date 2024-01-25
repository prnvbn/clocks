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

const (
	minMaxHeight = 5
)

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
		color, _ := pterm.NewGenericInteractiveSelect[Color]().
			WithMaxHeight(maxHeight()).
			WithOptions(Colors).
			WithDefaultText("Select a color for " + pterm.Bold.Sprint(z)).
			WithRenderSelectedOptionFunc(func(s string) string {
				return ColorFromString(s).ToStyle().
					Add(*pterm.Bold.ToStyle()).
					Sprintf("  %s\n", s)
			}).
			Show()

		heading, _ := pterm.DefaultInteractiveTextInput.
			WithMultiLine(false).
			WithDefaultText("Enter the display name for " + pterm.Bold.Sprint(z)).
			WithDefaultValue(z.City()).
			Show()

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
