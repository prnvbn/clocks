package cmd

import (
	"github.com/prnvbn/clocks/internal/ui"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:      "add [FUZZY_COUNTRY_NAME?]",
	Short:    "Add one or more clock",
	PostRunE: saveConfig,
	Args:     cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		oldNumClocks := len(cfg.ClockCfgs)
		searchTerm := ""
		if len(args) >= 1 {
			searchTerm = args[0]
		}
		clockCfgs := ui.SelectClocks(searchTerm)
		cfg.ClockCfgs.Add(clockCfgs...)

		newNumClocks := len(cfg.ClockCfgs)
		pterm.FgGreen.Println("Added", newNumClocks-oldNumClocks, "new clock(s)")
		printLayoutWarning()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
