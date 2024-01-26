package ui

import "github.com/pterm/pterm"

func EditClocks(ClockCfgs SortedClockConfigs) {
	clockCfg, _ := pterm.NewGenericInteractiveSelect[ClockConfig]().
		WithOptions(ClockCfgs).
		WithDefaultText("Select a clock to edit").
		WithMaxHeight(maxHeight()).
		WithFilter(false).
		WithRenderSelectedOptionFunc(func(s string) string {
			return ColorFromString(s).ToStyle().
				Add(*pterm.Bold.ToStyle()).
				Sprintf("> %s\n", s)
		}).
		Show()

	ClockCfgs.Remove(clockCfg)
	editClockConfig(&clockCfg)
	ClockCfgs.Add(clockCfg)
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
