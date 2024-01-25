package cmd

import (
	"github.com/prnvbn/clocks/internal/ui"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// layoutCmd represents the layout command
var layoutCmd = &cobra.Command{
	Use:                "layout",
	Short:              "Update clocks layout",
	PostRunE: saveConfig,
	Run: func(cmd *cobra.Command, args []string) {
		numClocks := len(cfg.ClockCfgs)
		layoutCfg := ui.SelectLayout(numClocks)
		cfg.Layout = layoutCfg

		pterm.FgGreen.Println("new layout set!")
		printLayoutWarning()
	},
}

func printLayoutWarning() bool {
	numClocks := len(cfg.ClockCfgs)
	if numClocks > cfg.Layout.MaxClocks() {
		pterm.FgYellow.Printfln(
			"Note: current layout only supports %d clock(s) but you have %d clocks added.\n"+
				"please remove some clocks using the 'remove' command or set a new layout using the 'layout' command.",
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
