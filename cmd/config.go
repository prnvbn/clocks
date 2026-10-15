package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "print the config file path and its contents",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Config file: %s\n\n", cfgPath)

		out, err := yaml.Marshal(cfg)
		fatal(err, "error marshaling config")

		fmt.Print(string(out))
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
