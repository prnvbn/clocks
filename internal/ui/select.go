package ui

import (
	"github.com/prnvbn/clocks/internal/tmz"
	"github.com/pterm/pterm"
)

func SelectZones(zones []tmz.Zone) (selected []tmz.Zone) {
	if len(zones) == 1 {
		selected = append(selected, zones[0])
	} else {
		tmzMenuHeading := pterm.ThemeDefault.PrimaryStyle.Sprint("Please select the timezones ") + pterm.ThemeDefault.SecondaryStyle.Sprint("[type to search]")
		selected, _ = pterm.NewGenericInteractiveMultiselect[tmz.Zone]().
			WithOptions(zones).
			WithDefaultText(tmzMenuHeading).
			WithClearAllEnabled(true).
			WithSelectAllEnabled(false).
			WithMaxHeight(maxHeight()).
			Show()
	}

	return
}

func SelectClockConfigs(cfgs []ClockConfig) (selected []ClockConfig) {
	if len(cfgs) == 1 {
		selected = append(selected, cfgs[0])
	} else {
		menuHeading := pterm.ThemeDefault.PrimaryStyle.Sprint("Please select the timezones ") + pterm.ThemeDefault.SecondaryStyle.Sprint("[type to search]")
		selected, _ = pterm.NewGenericInteractiveMultiselect[ClockConfig]().
			WithOptions(cfgs).
			WithDefaultText(menuHeading).
			WithClearAllEnabled(true).
			WithSelectAllEnabled(false).
			WithMaxHeight(maxHeight()).
			Show()
	}

	return
}
