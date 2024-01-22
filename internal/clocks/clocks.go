package clocks

import (
	"fmt"
	"time"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func Show(cfg Config) {
	clocks := make([][]string, 2, 3)

	fmt.Println(len(clocks), len(clocks[0]))
	// var clocks [][]string
	_ = clocks

	area, _ := pterm.DefaultArea.WithCenter(true).Start()

	defer area.Stop()

	leftText, _ := pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle(time.Now().Format("15:04"), pterm.FgCyan.ToStyle()),
	).Srender()
	rightText, _ := pterm.DefaultBigText.WithLetters(
		putils.LettersFromString(time.Now().Format("15:04")),
	).Srender()

	atr, _ := pterm.DefaultTable.
		WithData(
			[][]string{
				{leftText, rightText},
			},
		).
		WithSeparator("   "). // TODO? calc dynamically? with config?
		Srender()

	atr2, _ := pterm.DefaultTable.
		WithData(
			[][]string{
				{leftText, leftText, leftText},
			},
		).
		WithSeparator("   "). // TODO? calc dynamically? with config?
		Srender()

	area.Update(
		atr,
		atr2,
	)

	// for _, tmz := range cfg.Timzones {

	// }
}
