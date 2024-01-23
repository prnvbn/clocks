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
	verbose bool
	cfg     ui.AppConfig

	rootCmd = &cobra.Command{
		Use:   "clocks",
		Short: "display time across multiple timezones",
		Run:   run,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// TODO? graceful error handling
			yamlBytes, err := os.ReadFile(cfgPath)
			if err != nil {
				return err
			}

			err = yaml.Unmarshal(yamlBytes, &cfg)
			if err != nil {
				return err
			}

			err = cfg.Validate()
			if err != nil {
				return err
			}

			if verbose {
				zerolog.SetGlobalLevel(zerolog.DebugLevel)
			} else {
				zerolog.SetGlobalLevel(zerolog.Disabled)
			}

			log.Debug().Any("config", cfg).Msg("config loaded")
			if err != nil {
				return err
			}
			return nil
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			yamlBytes, err := yaml.Marshal(cfg)
			if err != nil {
				return err
			}
			// save yamlBytes to cfgPath
			err = os.WriteFile(cfgPath, yamlBytes, 0644)
			if err != nil {
				return err
			}

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
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgPath, "config", "c", defaultCfgFile, "Config file path, defaults to "+defaultCfgFile)
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enables additional logging")
}
