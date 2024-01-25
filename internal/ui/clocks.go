package ui

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

const separator = "   "

func ShowClocks(appCfg AppConfig) {

	area, _ := pterm.DefaultArea.Start()
	defer area.Stop()

	var str string
	if appCfg.Layout.CenterEachRow {
		str = showAsMultipleTables(appCfg)
	} else {
		str = showAsSingleTable(appCfg)
	}

	area.Update(str)
}

func showAsSingleTable(appCfg AppConfig) string {
	numRows := len(appCfg.Layout.RowSizes)
	clockIndex := 0

	tableData := make([][]string, 0, numRows)
	for i := 0; i < numRows; i++ {
		rowSize := appCfg.Layout.RowSizes[i]
		row := makeRow(rowSize, appCfg.ClockCfgs, &clockIndex)

		tableData = append(tableData, row)
	}

	clockTable, _ := pterm.DefaultTable.
		WithData(tableData).
		WithSeparator(separator).
		Srender()

	return centerRow(clockTable)
}

func showAsMultipleTables(appCfg AppConfig) string {
	numRows := len(appCfg.Layout.RowSizes)

	// each element in clockRows is a table srendern output that has N number of clocks
	// N is determined by layout.rowSizes
	//
	// The decision to use different table outputs for each row instead of
	// a single table was made because this allows us to center each row
	// independently
	clockRows := make([]string, numRows)

	clockIndex := 0
	for i, rowSize := range appCfg.Layout.RowSizes {

		row := makeRow(rowSize, appCfg.ClockCfgs, &clockIndex)
		clockTable, _ := pterm.DefaultTable.
			WithData([][]string{row}).
			WithSeparator("   ").
			Srender()

		clockRows[i] = clockTable
	}

	var ret string
	for _, cr := range clockRows {
		ret += centerRow(cr)
	}

	return ret
}

// makeRow creates a row of clocks of size ROW_SIZE
// The row is created from the clock configs in CLOCK_CFGS
// The row is created starting from the clock at index CLOCK_INDEX
// The clock index is incremented by the number of clocks in the row
func makeRow(rowSize int, clockCfgs SortedClockConfigs, clockIndex *int) []string {
	// only allocate capacity for the number of clocks in the row
	// this is because the number of clocks in this is less than
	// or equal to row size specified by the layout
	row := make([]string, 0, rowSize)

	for j := 0; j < rowSize; j++ {
		if *clockIndex == len(clockCfgs) {
			break
		}
		clockCfg := clockCfgs[*clockIndex]
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
		*clockIndex++
	}

	return row
}

func centerRow(s string) string {
	return pterm.DefaultCenter.Sprint(s)
}
