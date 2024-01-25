package cmd

import (
	"fmt"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all clocks",
	Run: func(cmd *cobra.Command, args []string) {

		bulletListItems := make([]pterm.BulletListItem, len(cfg.ClockCfgs))
		for i, clockCfg := range cfg.ClockCfgs {
			colorStyle := clockCfg.Color.ToStyle()
			bulletListItems[i] = pterm.BulletListItem{
				Level:       0,
				Text:        clockCfg.String(),
				TextStyle:   colorStyle,
				Bullet:      fmt.Sprintf("%02d.", i+1),
				BulletStyle: colorStyle,
			}
		}

		s, _ := pterm.DefaultBulletList.WithItems(bulletListItems).Srender()
		print(s)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
