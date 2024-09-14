package ui

import (
	"math"
	"os"
	"slices"

	"github.com/prnvbn/clocks/internal/tmz"
	"github.com/pterm/pterm"
	"golang.org/x/term"
)

const (
	minMaxHeight = 5
)

func SelectClocks(searchTerm string) []ClockConfig {
	cntries := slices.Collect(tmz.FilteredCountries(searchTerm))

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

	clockCfgs := make([]ClockConfig, len(selectedZones))

	for i, z := range selectedZones {
		clockCfg := ClockConfig{
			Zone:    z,
			Heading: z.City(),
		}
		editClockConfig(&clockCfg)

		clockCfgs[i] = clockCfg
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
