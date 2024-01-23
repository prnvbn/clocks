module github.com/prnvbn/clocks

go 1.21.3

require (
	github.com/pterm/pterm v0.12.75
	github.com/rs/zerolog v1.31.0
	github.com/spf13/cobra v1.8.0
	golang.org/x/exp v0.0.0-20220909182711-5c715a9e8561
	golang.org/x/term v0.16.0
	gopkg.in/yaml.v2 v2.2.4
	gopkg.in/yaml.v3 v3.0.1
)

require (
	atomicgo.dev/cursor v0.2.0 // indirect
	atomicgo.dev/keyboard v0.2.9 // indirect
	atomicgo.dev/schedule v0.1.0 // indirect
	github.com/containerd/console v1.0.3 // indirect
	github.com/gookit/color v1.5.4 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/lithammer/fuzzysearch v1.1.8 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/xo/terminfo v0.0.0-20220910002029-abceb7e1c41e // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)

// replace github.com/pterm/pterm => /home/prnvbn/opensource/pterm

// https://github.com/pterm/pterm/issues/618
// and other stuff...
replace github.com/pterm/pterm => github.com/prnvbn/pterm v0.0.7
