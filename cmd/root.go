package cmd

import (
	"os"

	"github.com/prnvbn/clocks/internal/ui"
	"github.com/pterm/pterm"
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
	live    bool

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
}

func fatal(err error, fmt string, args ...any) {
	if err != nil {
		print(pterm.Fatal.Sprintfln(fmt, args...))
		os.Exit(1)
	}
}
