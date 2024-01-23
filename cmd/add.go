package cmd

import (
	"math"
	"os"
	"slices"

	"github.com/prnvbn/clocks/internal/tmz"
	"github.com/prnvbn/clocks/internal/ui"
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
		ui.Black.String():   pterm.FgBlack.ToStyle(),
		ui.Blue.String():    pterm.FgRed.ToStyle(),
		ui.Cyan.String():    pterm.FgGreen.ToStyle(),
		ui.Green.String():   pterm.FgYellow.ToStyle(),
		ui.Magenta.String(): pterm.FgBlue.ToStyle(),
		ui.Red.String():     pterm.FgMagenta.ToStyle(),
		ui.White.String():   pterm.FgCyan.ToStyle(),
		ui.Yellow.String():  pterm.FgWhite.ToStyle(),
	}
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Allows you to add a timezone to your clocks",
	Run:   add,
}

func add(cmd *cobra.Command, args []string) {
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
		color, _ := pterm.NewGenericInteractiveSelect[ui.Color]().
			WithMaxHeight(maxHeight()).
			WithOptions(ui.Colors).
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
