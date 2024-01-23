/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/prnvbn/clocks/internal/ui"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"
)

// removeCmd represents the remove command
var (
	all       bool
	removeCmd = &cobra.Command{
		Use:   "remove",
		Short: "Remove one or more clock",

		Run: func(cmd *cobra.Command, args []string) {
			if all {
				if len(cfg.ClockCfgs) == 0 {
					pterm.Success.Println("No clocks to remove!")
					return
				}
				cfg.ClockCfgs = nil
				pterm.Success.Println("Removed all clocks")
				return
			}

			clockCfgs := ui.SelectClockConfigs(maps.Keys(cfg.ClockCfgs))
			cfg.ClockCfgs.Remove(clockCfgs...)
		},
	}
)

func init() {
	rootCmd.AddCommand(removeCmd)

	removeCmd.Flags().BoolVar(&all, "all", false, "remove all clocks")
}
