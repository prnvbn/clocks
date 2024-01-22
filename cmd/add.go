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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Allows you to add a timezone to your clocks",

	Run: func(cmd *cobra.Command, args []string) {
		cntries := maps.Keys(tmz.CountriesMap)
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

		var tmzs []string
		for _, cntry := range selectedCntries {
			cntryTMZs := tmz.CountriesMap[cntry]
			for _, tz := range cntryTMZs {
				tmzs = append(tmzs, tz.String())
			}
		}

		var selectedTMZs []string
		if len(tmzs) == 1 {
			selectedTMZs = append(selectedTMZs, tmzs[0])
		} else {
			tmzMenuHeading := pterm.ThemeDefault.PrimaryStyle.Sprint("Please select countries; you can select the timezones after this ") + pterm.ThemeDefault.SecondaryStyle.Sprint("[type to search]")
			selectedTMZs, _ = pterm.DefaultInteractiveMultiselect.
				WithOptions(tmzs).
				WithDefaultText(tmzMenuHeading).
				WithClearAllEnabled(true).
				WithSelectAllEnabled(false).
				WithMaxHeight(h).
				Show()
		}

		pterm.Success.Println(selectedTMZs)

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
