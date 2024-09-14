package tmz

import (
	"iter"

	"github.com/prnvbn/clocks/internal/match"
)

func FilteredCountries(searchTerm string) iter.Seq[string] {
	return func(yield func(string) bool) {
		for cntry := range CountryZonesMap {
			if match.Fuzzy(searchTerm, cntry) {
				if !yield(cntry) {
					return
				}
			}
		}
	}

}
