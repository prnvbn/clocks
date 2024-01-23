package cmd

import (
	"github.com/prnvbn/clocks/internal/ui"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Allows you to add a timezone to your clocks",
	Run: func(cmd *cobra.Command, args []string) {
		clockCfgs := ui.AddClocMenu()

		cfg.Timzones = append(cfg.Timzones, clockCfgs...)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
