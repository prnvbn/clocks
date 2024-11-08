package ui

import (
	"github.com/pterm/pterm"
	"github.com/rs/zerolog/log"
)

func EditClocks(clockCfgs ClockConfigs) {
	clockCfg, _ := pterm.NewGenericInteractiveSelect[ClockConfig]().
		WithOptions(clockCfgs.Slice()).
		WithDefaultText("Select a clock to edit").
		WithMaxHeight(maxHeight()).
		WithFilter(false).
		WithRenderSelectedOptionFunc(func(s string) string {
			return ColorFromString(s).ToStyle().
				Add(*pterm.Bold.ToStyle()).
				Sprintf("> %s\n", s)
		}).
		Show()

	clockCfgs.Remove(clockCfg)
	log.Debug().Msgf("Removed %s", clockCfg)

	editClockConfig(&clockCfg)
	clockCfgs.Add(clockCfg)
	log.Debug().Msgf("Added %s", clockCfg)
}

// editClockConfig propmts the user for clock config information.
// for zone z, the user is prompted for a heading and a color.
func editClockConfig(clockCfg *ClockConfig) {
	color, _ := pterm.NewGenericInteractiveSelect[Color]().
		WithMaxHeight(maxHeight()).
		WithOptions(Colors).
		WithDefaultText("Select a color for " + pterm.Bold.Sprint(clockCfg.Zone)).
		WithRenderSelectedOptionFunc(func(s string) string {
			return ColorFromString(s).ToStyle().
				Add(*pterm.Bold.ToStyle()).
				Sprintf("  %s\n", s)
		}).
		Show()
	clockCfg.Color = color

	heading, _ := pterm.DefaultInteractiveTextInput.
		WithMultiLine(false).
		WithDefaultText("Enter the display name for " + pterm.Bold.Sprint(clockCfg.Zone)).
		WithDefaultValue(clockCfg.Heading).
		Show()

	clockCfg.Heading = heading
}
