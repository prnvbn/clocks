package cmd

import (
	"github.com/prnvbn/clocks/internal/ui"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// layoutCmd represents the layout command
var layoutCmd = &cobra.Command{
	Use:      "layout",
	Short:    "Update clocks layout",
	PostRunE: saveConfig,
	Run: func(cmd *cobra.Command, args []string) {
		numClocks := len(cfg.ClockCfgs)
		layoutCfg := ui.SelectLayout(numClocks)
		cfg.Layout = layoutCfg

		pterm.FgGreen.Println("new layout set!")
		printLayoutWarning()
	},
}

// printLayoutWarning prints a warning if the current layout does not support all clocks
// returns true if a warning was printed
func printLayoutWarning() bool {
	if cfg.Layout.LayoutType != ui.Custom {
		return false
	}

	numClocks := len(cfg.ClockCfgs)
	if numClocks > cfg.Layout.MaxClocks() {
		pterm.FgYellow.Printfln(
			"NOTE: current layout only supports %d clock(s) but you have %d clocks added.\n"+
				"please set a new layout using 'clocks layout' or remove some clocks using 'clocks remove'.",
			cfg.Layout.MaxClocks(),
			numClocks,
		)
		return true
	}

	return false
}

func init() {
	rootCmd.AddCommand(layoutCmd)
}
