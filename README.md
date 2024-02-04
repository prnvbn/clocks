# clocks

clocks is a command line tool for displaying multiple clocks in your terminal. It allows you to configure multiple clocks and display them in a single view.

- Horizontal

  ![image](https://github.com/prnvbn/clocks/assets/55818107/2ace2664-7c58-4c30-b42a-e1b2cacdcd7f)


## Getting Started

### Installation

Run the following command to install clocks:

```bash
# TODO: curl command
```

or with wget:

```bash
# TODO: wget command
```

or with go:

```bash
as a go tool
```

TODO: Add installation instructions

This script installs the binary

### Adding a clock

Run `clocks add` and follow the prompts.

TODO: Add screenshot/video?

You can remove a clock by running `clocks remove` and following the prompts.
You can edit a clock by running `clocks edit` and following the prompts.

To list all clocks without the time, run `clocks list`.

## Additional Features

### Live mode

To keep the clocks running and updating in real time, run `clocks --live`. To persist this setting, run `clocks set --live`. This setting can be reversed by running `clocks unset --live`.

TODO: add screenshot/video?

### Seconds mode

To also see the seconds value, run `clocks --seconds`. To persist this setting, run `clocks set --seconds`. This setting can be reversed by running `clocks unset --seconds`.

### 12 hour mode

To display the time in 12 hour format, run `clocks --t12`. To persist this setting, run `clocks set --t12`. This setting can be reversed by running `clocks unset --t12`.

### Layouts
- Custom (Grid)

  ![image](https://github.com/prnvbn/clocks/assets/55818107/e0130fea-ffd8-4ea6-8edf-c086c9a4f176)

- Custom (Centered)
  
  ![image](https://github.com/prnvbn/clocks/assets/55818107/ab20d59a-b7a1-4691-b030-b3be31a8fe6a)

## Why

Sure, telling time in multiple places isn't rocket science. You could probably Google it or check your phone, however, I find this extremely annoying.

This annoyance is compounded when I am working on multiple projects simultaneously and monitoring different dashboards; thus leaving me with scarce screen real estate.

## Contributing

Feel free to open an issue or submit a pull request. I'm always open to suggestions and improvements :)

## TODO

[ ] CI/CD - release
[ ] Installation script
[ ] Update README TODOs
