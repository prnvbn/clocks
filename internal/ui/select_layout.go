package ui

import (
	"errors"
	"slices"
	"strconv"
	"strings"

	"github.com/pterm/pterm"
)

type LayoutType int

const (
	Horizontal LayoutType = iota
	Vertical
	Custom
)

var (
	Layouts = []LayoutType{Horizontal, Vertical, Custom}
	layouts = []string{
		"horizontal (all clocks in one row)",
		"vertical   (all clocks in one column)",
		"custom     (specify number of rows and clocks in each row)",
	}
)

func (l LayoutType) String() string {
	return layouts[l]
}

func SelectLayout(numClocks int) LayoutConfig {
	menuHeading := pterm.ThemeDefault.PrimaryStyle.Sprint("Please select a layout")
	selected, _ := pterm.NewGenericInteractiveSelect[LayoutType]().
		WithMaxHeight(maxHeight()).
		WithOptions(Layouts).
		WithFilter(false).
		WithRenderSelectedOptionFunc(func(s string) string {
			return pterm.ThemeDefault.SecondaryStyle.Sprintf("> %s\n", s)
		}).
		WithDefaultText(menuHeading).
		Show()

	switch selected {
	case Horizontal:
		return NewHorizontalLayout(numClocks)

	case Vertical:
		return NewVerticalLayout(numClocks)

	case Custom:
		pterm.Bold.Println("You currently have", numClocks, "clocks added, keep this in mind while setting a cutom layout")

		// In the custom case, first ask for number of rows they want
		r := repeat(func() (int, error) {
			rawNumRows, _ := pterm.DefaultInteractiveTextInput.
				WithDefaultText("Enter the number of rows").
				WithMultiLine(false).
				Show()

			return parseInt(rawNumRows)
		})

		rowSizes := make([]int, r)

		// Then ask for number of clocks in each row
		for i := 0; i < r; i++ {
			rowSizes[i] = repeat(func() (int, error) {
				rawNumClocks, _ := pterm.DefaultInteractiveTextInput.
					WithDefaultText("Enter the number of clocks in row " + strconv.Itoa(i+1)).
					WithMultiLine(false).
					Show()

				return strconv.Atoi(rawNumClocks)
			})
		}

		// Ask if they want to center each row individually
		centerEachRow := repeat(func() (bool, error) {
			defaultText := "Do you want to center each row individually?" +
				pterm.ThemeDefault.SecondaryStyle.Sprint(" [(y)es/(n)o]")
			ans, _ := pterm.DefaultInteractiveTextInput.
				WithDefaultText(defaultText).
				Show()

			if len(ans) == 0 {
				return false, errors.New("please enter yes or no")
			}

			ans = strings.ToLower(ans)
			if !slices.Contains([]string{"yes", "y", "no", "n"}, ans) {
				return false, errors.New("invalid option")
			}
			return ans[0] == 'y', nil
		})

		return LayoutConfig{
			RowSizes:      rowSizes,
			CenterEachRow: centerEachRow,
			LayoutType:    Custom,
		}
	default:
		panic("should never reach here")
	}
}

func parseInt(str string) (n int, err error) {
	n, err = strconv.Atoi(str)
	if errors.Is(err, strconv.ErrSyntax) {
		return 0, errors.New("please enter a valid number")
	}

	if n < 0 {
		return 0, errors.New("please enter a positive number")
	}

	return
}
func repeat[T any](askFunc func() (T, error)) T {
	for {
		val, err := askFunc()
		if err != nil {
			pterm.Error.Println(err)
			continue
		}
		return val
	}
}
