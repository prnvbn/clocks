# clocks

clocks is a command line tool for displaying multiple clocks in your terminal. It allows you to configure multiple clocks and display them in a single view. You can give different timezone clocks different colors to allow for easy distinction. You can also configure your preferred layout.

![image](https://github.com/prnvbn/clocks/assets/55818107/2ace2664-7c58-4c30-b42a-e1b2cacdcd7f)

## Installation

### Unix like systems

You can install the latest version of clocks by running the following command in your terminal. Remember to replace `<PLATFORM>` with your platform (linux-arm64, linux-amd64 or darwin).

```bash
curl -fsSL https://raw.githubusercontent.com/prnvbn/clocks/main/installer.sh | PLATFORM=<PLATFORM> bash
```

or using wget:

```bash
wget -qO - https://raw.githubusercontent.com/prnvbn/clocks/main/installer.sh |  PLATFORM=<PLATFORM> bash
```

Move the binary to a directory in your PATH. For e.g. `/usr/local/bin` on linux.

### Windows

Windows installation instructions are a WIP. In the meantime, you can download the latest release from the [releases page](https://github.com/prnvbn/clocks/releases)

## Getting Started

Run `clocks add` and follow the prompts to add a clock.
You can edit a clock at any time by running `clocks edit` and following the prompts.
The clocks can be displayed by running `clocks`.

You can remove a clock by running `clocks remove` and following the prompts.

To list all clocks without the time, run `clocks list`.

## Additional Features

### Layouts

Use the `clocks layout` comand to change the layout of the clocks.

- Custom (Grid)

  ![image](https://github.com/prnvbn/clocks/assets/55818107/e0130fea-ffd8-4ea6-8edf-c086c9a4f176)

- Custom (Centered)

  ![image](https://github.com/prnvbn/clocks/assets/55818107/ab20d59a-b7a1-4691-b030-b3be31a8fe6a)

### Seconds mode

To also see the seconds value, run `clocks --seconds`. To persist this setting, run `clocks set --seconds`. This setting can be reversed by running `clocks unset --seconds`.
![image](https://github.com/prnvbn/clocks/assets/55818107/94cef848-952a-4526-b4ee-2193e3219100)

### 12 hour mode

To display the time in 12 hour format, run `clocks --t12`. To persist this setting, run `clocks set --t12`. This setting can be reversed by running `clocks unset --t12`.
![image](https://github.com/prnvbn/clocks/assets/55818107/13a160a8-c442-477c-be01-7b58df9e99b2)

Note that additional work to display AM/PM on the clock face as well

### Live mode

To keep the clocks running and updating in real time, run `clocks --live`. To persist this setting, run `clocks set --live`. This setting can be reversed by running `clocks unset --live`.

## Configuration

clocks is configured by a YAML file that will be auto-generated on first run. By default, the config file is assumed to exist on an [XDG-compliant](https://en.wikipedia.org/wiki/Freedesktop.org) configuration path like `~/.config/clocks/config.yaml`. If you would like to store it elsewhere, you may export a `CLOCKS_CONFIG_PATH` environment variable that specifies its path:

```bash
export CLOCKS_CONFIG_PATH="<NEW_CONFIG_PATH>"
```

## Why

Sure, telling time in multiple places isn't rocket science. You could probably Google it or check your phone, however, I find this extremely annoying.

This annoyance is compounded when I am working on multiple projects simultaneously and monitoring different dashboards; thus leaving me with scarce screen real estate.

## Contributing

Feel free to open an issue or submit a pull request. I'm always open to suggestions and improvements :)
