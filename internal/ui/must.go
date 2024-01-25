package ui

import "fmt"

func must(err error) {
	if err != nil {
		panic(fmt.Errorf("pterm error: %w", err))
	}
}
