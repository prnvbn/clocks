/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all clocks",
	Run: func(cmd *cobra.Command, args []string) {

		bulletListItems := make([]pterm.BulletListItem, len(cfg.ClockCfgs))
		for i, clockCfg := range maps.Keys(cfg.ClockCfgs) {
			colorStyle := clockCfg.Color.ToStyle()
			bulletListItems[i] = pterm.BulletListItem{
				Level:       0,
				Text:        clockCfg.String(),
				TextStyle:   colorStyle,
				Bullet:      fmt.Sprintf("%02d.", i+1),
				BulletStyle: colorStyle,
			}
		}

		// Have to use Srender() instead of Render() because of a bug(?) in pterm
		// extra newline is added when using Render()
		s, _ := pterm.DefaultBulletList.WithItems(bulletListItems).Srender()
		print(s)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
