# clocks

clocks is a command line tool for displaying multiple clocks in your terminal. It allows you to configure multiple clocks and display them in a single view.

## Getting Started

### Installation

TODO: Add installation instructions

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

To display the time in 12 hour format, run `clocks --12`. To persist this setting, run `clocks set --12`. This setting can be reversed by running `clocks unset --12`.

## Why

Sure, telling time in multiple places isn't rocket science. You could probably Google it or check your phone. However, that is extremely annoying (for me at the very least).

This annoyance is compounded when I am working on multiple projects simultaneously and monitoring different dashboards; thus leaving me with scarce screen real estate.

## Contributing

Feel free to open an issue or submit a pull request. I'm always open to suggestions and improvements :)

## TODO

[ ] CI/CD - release
[ ] Installation script
[ ] Update README TODOs
