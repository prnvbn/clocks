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
	unsetCmd = &cobra.Command{
		Use:      "unset",
		Short:    "Permanently unset flags for the 'clocks' command",
		PostRunE: saveConfig,
		Run: func(cmd *cobra.Command, args []string) {
			if !liveChanged && !secondsChanged && !twelveHrChanged {
				pterm.FgRed.Println("please pass a valid flag")
				return
			}

			if liveChanged {
				cfg.Live = false
			}
			if secondsChanged {
				cfg.Seconds = false
			}
			if twelveHrChanged {
				cfg.TwelveHour = false
			}

			pterm.FgGreen.Println("config updated accordingly!")
		},
	}
)

func init() {
	rootCmd.AddCommand(unsetCmd)

	addFlags(unsetCmd)
}
