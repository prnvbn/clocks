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
			if !live && !seconds {
				pterm.FgRed.Println("please pass a valid flag")
				return
			}

			cfg.Live = false
			cfg.Seconds = false

			pterm.FgGreen.Println("config updated accordingly!")
		},
	}
)

func init() {
	rootCmd.AddCommand(unsetCmd)

	unsetCmd.Flags().BoolVarP(&live, "live", "l", false, "sets the live flag to true for all subsequent calls")
	unsetCmd.Flags().BoolVarP(&seconds, "seconds", "s", false, "sets the seconds flag to true for all subsequent calls")
}
