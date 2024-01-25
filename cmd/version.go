package cmd

import (
	_ "embed"
	"fmt"

	"github.com/spf13/cobra"
)

//go:embed version.txt
var version string

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "shows the tool version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("v%s", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
