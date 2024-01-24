package cmd

import (
	"github.com/prnvbn/clocks/internal/ui"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add one or more clock",
	Run: func(cmd *cobra.Command, args []string) {
		oldNumClocks := len(cfg.ClockCfgs)
		clockCfgs := ui.SelectClocks()
		cfg.ClockCfgs.Add(clockCfgs...)

		newNumClocks := len(cfg.ClockCfgs)
		pterm.FgGreen.Println("Added", newNumClocks-oldNumClocks, "new clock(s)")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
