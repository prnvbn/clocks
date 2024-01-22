/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/prnvbn/clocks/internal/clocks"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

const (
	defaultCfgFile = "~/.clocks.yaml"
)

var (
	// flags
	cfgPath string

	rootCmd = &cobra.Command{
		Use:   "clocks",
		Short: "display time across multiple timezones",
		Run:   run,
	}
)

func run(cmd *cobra.Command, args []string) {

	yamlBytes, err := os.ReadFile(cfgPath)
	must(err)

	var cfg clocks.Config
	err = yaml.Unmarshal(yamlBytes, &cfg)
	must(err)

	// TODO? validate config
	err = cfg.Validate()
	must(err)

	// TODO: add pretty printing for config when -v flag is passed
	fmt.Println(cfg)

	clocks.Show(cfg)

	// leftText, _ := pterm.DefaultBigText.WithLetters(
	// 	putils.LettersFromString(time.Now().Format("15:04")),
	// ).Srender()
	// rightText, _ := pterm.DefaultBigText.WithLetters(
	// 	putils.LettersFromString(time.Now().Format("15:04")),
	// ).Srender()

	// area, _ := pterm.DefaultArea.WithCenter().Start()
	// defer area.Stop()

	// for i := 0; i < 5; i++ {
	// 	atr, _ := pterm.DefaultTable.
	// 		WithData(
	// 			// TODO: getGrid(gridLayout)
	// 			[][]string{
	// 				{leftText, rightText},
	// 				{leftText},
	// 			},
	// 		).
	// 		WithSeparator("                "). // TODO? calc dynamically? with config?
	// 		Srender()

	// 	area.Update(atr)
	// }

}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgPath, "config", "c", defaultCfgFile, "Config file path, defaults to "+defaultCfgFile)
}

// TODO! graceful handling
func must(err error) {
	if err != nil {
		panic(err)
	}
}
