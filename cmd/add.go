package cmd

import (
	"math"
	"os"
	"slices"

	"github.com/prnvbn/clocks/internal/clocks"
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
	colorToStyle = map[string]*pterm.Style{
		clocks.Black.String():   pterm.FgBlack.ToStyle(),
		clocks.Blue.String():    pterm.FgRed.ToStyle(),
		clocks.Cyan.String():    pterm.FgGreen.ToStyle(),
		clocks.Green.String():   pterm.FgYellow.ToStyle(),
		clocks.Magenta.String(): pterm.FgBlue.ToStyle(),
		clocks.Red.String():     pterm.FgMagenta.ToStyle(),
		clocks.White.String():   pterm.FgCyan.ToStyle(),
		clocks.Yellow.String():  pterm.FgWhite.ToStyle(),
	}
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Allows you to add a timezone to your clocks",

	Run: func(cmd *cobra.Command, args []string) {
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
				WithMaxHeight(maxHeight()).
				Show()
		}

		for _, z := range selectedZones {
			// TODO? show sample number in an area next to the select menu
			color, _ := pterm.NewGenericInteractiveSelect[clocks.Color]().
				WithMaxHeight(maxHeight()).
				WithOptions(clocks.Colors).
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
			pterm.Success.Println(color.String() + " selected for " + z.String() + " with name " + result)
		}
	},
}

func maxHeight() int {
	_, h, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		h = minMaxHeight
	}
	return int(math.Max(1, float64(h-5)))
}

func init() {
	rootCmd.AddCommand(addCmd)
}
