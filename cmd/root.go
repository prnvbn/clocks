/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/prnvbn/clocks/internal/ui"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

const (
	defaultCfgFile = "~/.clocks.yaml"
)

var (
	// flags
	cfgPath string
	debug   bool
	cfg     ui.AppConfig

	rootCmd = &cobra.Command{
		Use:   "clocks",
		Short: "display time across multiple timezones",
		Run:   run,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// TODO? graceful error handling
			yamlBytes, err := os.ReadFile(cfgPath)
			must(err)

			err = yaml.Unmarshal(yamlBytes, &cfg)
			must(err)
			err = cfg.Validate()
			must(err)

			if debug {
				zerolog.SetGlobalLevel(zerolog.DebugLevel)
			} else {
				zerolog.SetGlobalLevel(zerolog.Disabled)
			}

			log.Debug().Interface("config", cfg).Msg("config loaded")
			must(err)

			return nil
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			yamlBytes, err := yaml.Marshal(cfg)
			must(err)

			// save yamlBytes to cfgPath
			err = os.WriteFile(cfgPath, yamlBytes, 0644)
			must(err)

			log.Debug().Any("config", cfg).Msg("config saved")
			return nil
		},
	}
)

func run(cmd *cobra.Command, args []string) {
	ui.ShowClocks(cfg)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		// TODO? graceful error handling
		// add recover that gives an option to raise a GH issue
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgPath, "config", "c", defaultCfgFile, "Config file path, defaults to "+defaultCfgFile)
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "enables debug logging")
	_ = rootCmd.PersistentFlags().MarkHidden("debug")
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
