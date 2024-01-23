package cmd

import (
	"math"
	"os"
	"slices"

	"github.com/prnvbn/clocks/internal/tmz"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"
	"golang.org/x/term"
)

const (
	minMaxHeight = 5
)

var (
	colors = []string{
		"black",
		"red",
		"green",
		"yellow",
		"blue",
		"magenta",
		"cyan",
		"white",
	}

	colorToStyle = map[string]*pterm.Style{
		"black":   pterm.FgBlack.ToStyle(),
		"red":     pterm.FgRed.ToStyle(),
		"green":   pterm.FgGreen.ToStyle(),
		"yellow":  pterm.FgYellow.ToStyle(),
		"blue":    pterm.FgBlue.ToStyle(),
		"magenta": pterm.FgMagenta.ToStyle(),
		"cyan":    pterm.FgCyan.ToStyle(),
		"white":   pterm.FgWhite.ToStyle(),
	}
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Allows you to add a timezone to your clocks",

	Run: func(cmd *cobra.Command, args []string) {
		cntries := maps.Keys(tmz.CountryZonesMap)
		slices.Sort(cntries)

		_, h, err := term.GetSize(int(os.Stdout.Fd()))
		if err != nil {
			h = minMaxHeight
		}
		h = int(math.Max(1, float64(h-5)))

		cntryMenuHeading := pterm.ThemeDefault.PrimaryStyle.Sprint("Please select countries; you can select the timezones after this ") + pterm.ThemeDefault.SecondaryStyle.Sprint("[type to search]")
		selectedCntries, _ := pterm.DefaultInteractiveMultiselect.
			WithOptions(cntries).
			WithDefaultText(cntryMenuHeading).
			WithClearAllEnabled(true).
			WithSelectAllEnabled(false).
			WithMaxHeight(h).
			Show()

		var zones []tmz.Zone
		for _, cntry := range selectedCntries {
			cntryToZones := tmz.CountryZonesMap[cntry]
			zones = append(zones, cntryToZones...)
		}

		var selectedZones []tmz.Zone
		if len(zones) == 1 {
			selectedZones = append(selectedZones, zones[0])
		} else {
			tmzMenuHeading := pterm.ThemeDefault.PrimaryStyle.Sprint("Please select the timezones ") + pterm.ThemeDefault.SecondaryStyle.Sprint("[type to search]")
			selectedZones, _ = pterm.NewGenericInteractiveMultiselect[tmz.Zone]().
				WithOptions(zones).
				WithDefaultText(tmzMenuHeading).
				WithClearAllEnabled(true).
				WithSelectAllEnabled(false).
				WithMaxHeight(h).
				Show()
		}

		for _, z := range selectedZones {
			// TODO? show sample number in an area next to the select menu
			color, _ := pterm.DefaultInteractiveSelect.
				WithMaxHeight(h).
				WithOptions(colors).
				WithDefaultText("Select a color for " + pterm.Bold.Sprint(z)).
				WithRenderSelectedOptionFunc(func(s string) string {
					return colorToStyle[s].
						Add(*pterm.Bold.ToStyle()).
						Sprintf("  %s\n", s)
				}).
				Show()

			textInput := pterm.DefaultInteractiveTextInput.
				WithMultiLine(false).
				WithDefaultText("Enter a name for " + pterm.Bold.Sprint(z)).
				WithDefaultValue(z.City())

			result, _ := textInput.Show()

			pterm.Println()
			pterm.Info.Printfln("You answered: %s", result)
			pterm.Success.Println(color + " selected for " + z.String() + " with name " + result)
		}
	},
}

func init() {
	slices.Sort(colors)
	rootCmd.AddCommand(addCmd)
}
