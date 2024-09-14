package match

import (
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

func Fuzzy(searchTerm string, toMatch string) bool {
	return fuzzy.Match(strings.ToLower(searchTerm), strings.ToLower(toMatch))
}
