package ui

import (
	"strings"
	"time"

	"github.com/mattn/go-runewidth"
	"github.com/pkg/errors"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

const separator = "   "

func ShowClocks(appCfg AppConfig) {
	area, _ := pterm.DefaultArea.Start()
	defer func() {
		if err := area.Stop(); err != nil {
			panic(errors.Wrap(err, "failed to stop pterm area"))
		}
	}()

	for {
		var str string
		if appCfg.Layout.CenterEachRow {
			str = showAsMultipleTables(appCfg)
		} else {
			str = showAsSingleTable(appCfg)
		}
		area.Update(str)

		if !appCfg.Live {
			break
		}

		time.Sleep(time.Second)
	}
}

func showAsSingleTable(appCfg AppConfig) string {
	numRows := len(appCfg.Layout.RowSizes)
	clockIndex := 0

	tableData := make([][]string, 0, numRows)
	for i := 0; i < numRows; i++ {
		rowSize := appCfg.Layout.RowSizes[i]
		row := makeRow(rowSize, appCfg, &clockIndex)

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

		row := makeRow(rowSize, appCfg, &clockIndex)
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
// The row is created from the clock configs in APP_CFG
// The row is created starting from the clock at index CLOCK_INDEX
// The clock index is incremented by the number of clocks in the row
func makeRow(rowSize int, appCfg AppConfig, clockIndex *int) []string {
	clockCfgs := appCfg.ClockCfgs
	dateFmt := "15:04"
	if appCfg.TwelveHour {
		dateFmt = "3:04"
	}
	if appCfg.Seconds {
		dateFmt += ":05"
	}

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
		now := clockCfg.Zone.Now()
		style := clockCfg.Color.ToStyle()
		styledTime, _ := pterm.DefaultBigText.WithLetters(
			putils.LettersFromStringWithStyle(
				now.Format(dateFmt),
				style,
			),
		).Srender()
		heading := style.Sprint(clockCfg.Heading)
		if appCfg.TwelveHour {
			clockWidth := 0
			for _, line := range strings.Split(pterm.RemoveColorFromString(styledTime), "\n") {
				clockWidth = max(clockWidth, runewidth.StringWidth(strings.TrimRight(line, " ")))
			}
			padding := max(1, clockWidth-runewidth.StringWidth(clockCfg.Heading)-2)
			heading += strings.Repeat(" ", padding) + style.Sprint(now.Format("pm"))
		}

		row = append(row, heading+"\n"+styledTime)
		*clockIndex++
	}

	return row
}

func centerRow(s string) string {
	return pterm.DefaultCenter.Sprint(s)
}
