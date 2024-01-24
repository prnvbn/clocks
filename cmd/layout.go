package cmd

import (
	"github.com/prnvbn/clocks/internal/ui"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// layoutCmd represents the layout command
var layoutCmd = &cobra.Command{
	Use:   "layout",
	Short: "Update clocks layout",

	Run: func(cmd *cobra.Command, args []string) {
		numClocks := len(cfg.ClockCfgs)
		layoutCfg := ui.SelectLayout(numClocks)
		cfg.Layout = layoutCfg

		maxLayouts := layoutCfg.MaxClocks()

		pterm.FgGreen.Println("new layout set!")
		if numClocks > maxLayouts {
			pterm.FgYellow.Printfln(
				"Note: current layout only supports %d clock(s) but you have %d clocks added.\n"+
					"please remove some clocks using the 'remove' command or set a new layout",
				maxLayouts,
				numClocks,
			)
		}
	},
}

func init() {
	rootCmd.AddCommand(layoutCmd)
}
