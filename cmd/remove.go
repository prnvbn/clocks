package cmd

import (
	"github.com/prnvbn/clocks/internal/ui"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var (
	all       bool
	removeCmd = &cobra.Command{
		Use:      "remove",
		Short:    "Remove one or more clock",
		PostRunE: saveConfig,
		Run: func(cmd *cobra.Command, args []string) {
			if all {
				if len(cfg.ClockCfgs) == 0 {
					pterm.FgYellow.Println("No clocks to remove!")
					return
				}
				cfg.ClockCfgs = nil
				pterm.FgGreen.Println("Removed all clocks")
				return
			}

			clockCfgs := ui.SelectClockConfigs(cfg.ClockCfgs)
			cfg.ClockCfgs.Remove(clockCfgs...)
			pterm.FgGreen.Println("Removed", len(clockCfgs), "clock(s)")
		},
	}
)

func init() {
	rootCmd.AddCommand(removeCmd)

	removeCmd.Flags().BoolVar(&all, "all", false, "remove all clocks")
}
