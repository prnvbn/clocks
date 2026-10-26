/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var (
	setCmd = &cobra.Command{
		Use:      "set",
		Short:    "Permanently set flags for the 'clocks' command",
		PostRunE: saveConfig,
		Run: func(cmd *cobra.Command, args []string) {
			if !liveChanged && !secondsChanged && !twelveHrChanged {
				pterm.FgRed.Println("please pass a valid flag")
				return
			}

			if liveChanged {
				cfg.Live = live
			}
			if secondsChanged {
				cfg.Seconds = seconds
			}
			if twelveHrChanged {
				cfg.TwelveHour = twelveHr
			}

			pterm.FgGreen.Println("config updated accordingly!")
		},
	}
)

func init() {
	rootCmd.AddCommand(setCmd)

	addFlags(setCmd)
}
