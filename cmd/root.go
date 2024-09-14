package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	rdebug "runtime/debug"

	"github.com/adrg/xdg"
	"github.com/prnvbn/clocks/internal/ui"
	"github.com/pterm/pterm"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

const (
	REPORT_LINK = "https://github.com/prnvbn/clocks/issues/new"
)

var (
	// flags
	cfgPath  string
	debug    bool
	live     bool
	seconds  bool
	twelveHr bool

	cfg     ui.AppConfig
	rootCmd = &cobra.Command{
		Use:               "clocks [FUZZY_CLOCK_NAME?]",
		Short:             "A tool to display time across multiple timezones.",
		PersistentPreRunE: loadConfig,
		Args:              cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if clocksAbsent() {
				return
			}

			if printLayoutWarning() {
				pterm.FgRed.Println("Can not display all clocks with current layout.")
				return
			}

			if live {
				cfg.Live = true
			}
			if seconds {
				cfg.Seconds = true
			}

			if twelveHr {
				cfg.TwelveHour = true
			}
		
			if len(args) >= 1 {
				// when a @search term is passed, set layout to horizontal
				// so that all clocks are displayed in one row
				// clocks are filtered by the search term 
				// fuzzy search is used to match the search term
				searchTerm := args[0]
				cfg.Layout.LayoutType = ui.Horizontal

				n := 0
				cfg.ClockCfgs, n = cfg.ClockCfgs.Filter(searchTerm)
				if n == 0 {
					pterm.FgYellow.Println("No clocks match the search term:", searchTerm)
					return
				}
			}

			if len(args) >= 1 {
				// when a @search term is passed, set layout to horizontal
				// so that all clocks are displayed in one row
				// clocks are filtered by the search term
				// fuzzy search is used to match the search term
				searchTerm := args[0]
				cfg.Layout.LayoutType = ui.Horizontal

				n := 0
				cfg.ClockCfgs, n = cfg.ClockCfgs.Filter(searchTerm)
				if n == 0 {
					pterm.FgYellow.Println("No clocks match the search term:", searchTerm)
					return
				}
			}


			ui.ShowClocks(cfg)
		},
	}
)

// loadConfig loads the config from cfgPath
// If the config does not exist, it creates a new one
func loadConfig(cmd *cobra.Command, args []string) error {
	yamlBytes, err := os.ReadFile(cfgPath)
	if err != nil {
		if os.IsNotExist(err) {
			// create config file
			err = os.MkdirAll(filepath.Dir(cfgPath), 0755)
			fatal(err, "error creating config directory at %s", filepath.Dir(cfgPath))

			err = saveConfig(cmd, args)
			fatal(err, "error saving config file at %s", cfgPath)
		}
		return err
	}
	fatal(err, "Config file %s does not exist!", cfgPath)

	err = yaml.Unmarshal(yamlBytes, &cfg)
	fatal(err, "YAML Config file at %s is malformed", cfgPath)

	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.Disabled)
	}

	// Update layout for vertical and horizontal layouts
	// this is done so that once the user sets their layout as vertical or horizontal
	// it stays that way even if they add or remove clocks
	if cfg.Layout.LayoutType == ui.Horizontal {
		cfg.Layout = ui.NewHorizontalLayout(len(cfg.ClockCfgs))
	} else if cfg.Layout.LayoutType == ui.Vertical {
		cfg.Layout = ui.NewVerticalLayout(len(cfg.ClockCfgs))
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
}

func init() {

	var ok bool
	cfgPath, ok = os.LookupEnv("CLOCKS_CONFIG_PATH")
	if !ok {
		// default to $XDG_CONFIG_HOME/clocks/config.yaml
		cfgPath = xdg.ConfigHome + "/clocks/config.yaml"
	}

	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "enables debug logging")
	_ = rootCmd.PersistentFlags().MarkHidden("debug")
	_ = rootCmd.PersistentFlags().MarkHidden("config")

	addFlags(rootCmd)
}

func addFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&live, "live", "l", false, "keeps clocks on screen")
	cmd.Flags().BoolVarP(&seconds, "seconds", "s", false, "shows seconds as well")
	cmd.Flags().BoolVar(&twelveHr, "t12", false, "print dates in 12 hour format")
}

func fatal(err error, fmt string, args ...any) {
	if err != nil {
		print(pterm.Fatal.Sprintfln(fmt, args...))
		os.Exit(1)
	}
}
