# clocks

clocks is a command line tool for displaying multiple clocks in your terminal. It allows you to configure multiple clocks and display them in a single view. You can give different timezone clocks different colors to allow for easy distinction. You can also configure your preferred layout.

![Terminal demo showing clocks, filtering by city, and live updates with seconds](demo/clocks-demo.gif)

## Installation

### Linux or MacOS

You can install the latest version of clocks by running the following command in your terminal.

```bash
bash -c "$(curl -fsSL https://raw.githubusercontent.com/prnvbn/clocks/main/installer.sh)"
```

Move the binary to a directory in your PATH. For e.g. `/usr/local/bin` on linux.

### via `homebrew`

You can install `clocks` using the [prnvbn/homebrew-tap](https://github.com/prnvbn/homebrew-tap).

```bash
brew install prnvbn/tap/clocks
```

### via `go install`

```bash
go install github.com/prnvbn/clocks@latest
```

#### Enabling Command Autocompletion

To enable autocomplete, add the following to your `.bashrc` or `.bash_profile` file:

```bash
# you can also generate completions for zsh and fish shells by replacing bash with zsh or fish
source <(clocks completion bash)
```

> NOTE: dont forget to restart your terminal or run `source ~/.bashrc`

If you want use an alias for clocks, you can add the following to your `.bashrc` or `.bash_profile` file:

```bash
alias c=clocks
source <(clocks completion bash | sed 's/clocks/c/g')
```

## Getting Started

Run `clocks add` and follow the prompts to add a clock.
You can edit a clock at any time by running `clocks edit` and following the prompts.
The clocks can be displayed by running `clocks`.

You can remove a clock by running `clocks remove` and following the prompts.

To list all clocks without the time, run `clocks list`.

## Additional Features

### Layouts

Use the `clocks layout` command to change the layout of the clocks.
Use `clocks --horizontal` (or `clocks -H`) to display all clocks in one row for a single run, or `clocks --vertical` (or `clocks -V`) to display them in one column. These flags override the configured layout without changing it.

- Custom (Grid)

  ![Three clocks in a grid with aligned columns](demo/layout-grid.png)

- Custom (Centered)

  ![Three clocks with each row centered independently](demo/layout-centered.png)

### Seconds mode

To also see the seconds value, run `clocks --seconds`. To persist this setting, run `clocks set --seconds`. This setting can be reversed by running `clocks unset --seconds`.
![Clocks showing seconds and AM/PM indicators](demo/seconds.png)

### 24 hour mode

To use 24 hour time, run `clocks unset --t12`, then run `clocks`. This saves your preference and removes the AM/PM indicators.

New configs use 12 hour time by default. Run `clocks set --t12` to switch back, or `clocks --t12` to use 12 hour time for one run.

![London and New York clocks showing 24 hour time](demo/twenty-four-hour.png)

### Live mode

To keep the clocks running and updating in real time, run `clocks --live`. To persist this setting, run `clocks set --live`. This setting can be reversed by running `clocks unset --live`.

## Configuration

clocks is configured by a YAML file that will be auto-generated on first run with `twelveHour: true`. Existing configs keep their settings. By default, the config file is stored at `$XDG_CONFIG_HOME/clocks/config.yaml` when `XDG_CONFIG_HOME` is set, or `~/.config/clocks/config.yaml` on Linux and `~/Library/Application Support/clocks/config.yaml` on macOS. Run `clocks config` to see its path and contents. If you would like to store it elsewhere, you may export a `CLOCKS_CONFIG_PATH` environment variable that specifies its path:

```bash
export CLOCKS_CONFIG_PATH="<NEW_CONFIG_PATH>"
```

## Why

Sure, telling time in multiple places isn't rocket science. You could probably Google it or check your phone, however, I find this extremely annoying.

This annoyance is compounded when I am working on multiple projects simultaneously and monitoring different dashboards; thus leaving me with scarce screen real estate.

## Contributing

Feel free to open an issue or submit a pull request. I'm always open to suggestions and improvements :)
