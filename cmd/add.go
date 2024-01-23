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
		clockCfgs := ui.SelectClocks()
		cfg.ClockCfgs.Add(clockCfgs...)
		pterm.Success.Println("Added", len(clockCfgs), "clock(s)")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
