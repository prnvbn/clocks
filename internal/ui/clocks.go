package ui

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func ShowClocks(appCfg AppConfig) {

	numRows := len(appCfg.Layout.RowSizes)

	// each element in clockRows is a table output that has N number of clocks
	// N is determined by layout.rowSizes
	//
	// The decision to use different table outputs for each row instead of
	// a single table was made because this allows us to center each row
	// independently if needed
	clockRows := make([]string, numRows)

	area, _ := pterm.DefaultArea.
		WithCenter(appCfg.Layout.CenterEachRow).
		Start()
	defer area.Stop()

	clockIndex := 0

	for i, rowSize := range appCfg.Layout.RowSizes {
		// only allocate capacity for the number of clocks in the row
		// this is because the number of clocks in this is less than
		// or equal to row size specified by the layout
		row := make([]string, 0, rowSize)
		for j := 0; j < rowSize; j++ {
			if clockIndex >= len(appCfg.ClockCfgs) {
				break
			}
			clockCfg := appCfg.ClockCfgs[clockIndex]

			// get time based on zone
			zone := clockCfg.Zone
			style := clockCfg.Color.ToStyle()
			styledTime, _ := pterm.DefaultBigText.WithLetters(
				putils.LettersFromStringWithStyle(
					zone.Now().Format("15:04"), //TODO? make format configurable OR add a * in 12 hr format to indicate PM
					style,
				),
			).Srender()
			heading := style.Sprint(clockCfg.Heading)

			row = append(row, heading+"\n"+styledTime)
			clockIndex++
		}

		clockTable, _ := pterm.DefaultTable.
			WithData([][]string{row}).
			WithSeparator("   ").
			Srender()

		clockRows[i] = clockTable
	}

	// need to cast to []any because Update takes a variadic of []any
	castedClockRows := make([]any, len(clockRows))
	for i, row := range clockRows {
		castedClockRows[i] = row
	}

	area.Update(castedClockRows...)
}
