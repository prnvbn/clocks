package cmd

import (
	"fmt"
	"os"
	rdebug "runtime/debug"

	"github.com/prnvbn/clocks/internal/ui"
	"github.com/pterm/pterm"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

const (
	defaultCfgFile = "~/.clocks.yaml"
	REPORT_LINK    = "https://github.com/prnvbn/clocks/issues/new"
)

var (
	// flags
	cfgPath string
	debug   bool
	live    bool
	seconds bool

	cfg     ui.AppConfig
	rootCmd = &cobra.Command{
		Use:               "clocks",
		Short:             "display time across multiple timezones",
		PersistentPreRunE: prerun,
		Run: func(cmd *cobra.Command, args []string) {
			numClocks := len(cfg.ClockCfgs)
			if numClocks == 0 {
				pterm.FgYellow.Println("No clocks to display!")
				pterm.FgBlue.Println("HINT: Use 'clocks add' to add a clock")
				return
			}

			if printLayoutWarning() {
				return
			}

			if live {
				cfg.Live = true
			}
			if seconds {
				cfg.Seconds = true
			}

			ui.ShowClocks(cfg)
		},
	}
)

func prerun(cmd *cobra.Command, args []string) error {
	yamlBytes, err := os.ReadFile(cfgPath)
	fatal(err, "Config file %s does not exist!", cfgPath)

	err = yaml.Unmarshal(yamlBytes, &cfg)
	fatal(err, "YAML Config file at %s is malformed", cfgPath)

	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.Disabled)
	}

	log.Debug().Interface("config", cfg).Msg("config loaded")

	return nil
}

func saveConfig(cmd *cobra.Command, args []string) error {
	yamlBytes, err := yaml.Marshal(cfg)
	fatal(err, "YAML config file at %s is malformed", cfgPath)

	// save yamlBytes to cfgPath
	err = os.WriteFile(cfgPath, yamlBytes, 0644)
	fatal(err, "error saving YAML file to %s", cfgPath)

	log.Debug().Any("config", cfg).Msg("config saved")
	return nil
}

func Execute() {

	defer func() {
		if r := recover(); r != nil {
			pterm.BgRed.Println("Seems like you have discored a bug! Please report the issue at:", REPORT_LINK)
			pterm.BgRed.Println("Please attach the below stack trace as part of the bug report")

			fmt.Println("==================STACK-TRACE=====================")
			fmt.Println("Stack trace:", string(rdebug.Stack()))
			fmt.Println("=====================================================")

			pterm.BgRed.Println("Seems like you have discored a bug! Please report the issue at:", REPORT_LINK)
			pterm.BgRed.Println("Please attach the above stack trace as part of the bug report")
		}
	}()

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	// TODO? graceful error handling
	// add panic recover that gives an option to raise a GH issue
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgPath, "config", "c", defaultCfgFile, "Config file path, defaults to "+defaultCfgFile)
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "enables debug logging")
	_ = rootCmd.PersistentFlags().MarkHidden("debug")

	rootCmd.Flags().BoolVarP(&live, "live", "l", false, "keeps clocks on screen")
	rootCmd.Flags().BoolVarP(&seconds, "seconds", "s", false, "shows seconds as well")
}

func fatal(err error, fmt string, args ...any) {
	if err != nil {
		print(pterm.Fatal.Sprintfln(fmt, args...))
		os.Exit(1)
	}
}
