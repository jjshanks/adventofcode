package challenge

import (
	"sort"
	"strings"
)

type Observation struct {
	Patterns []string
	Digits   []string
}

func parse(instr string) []*Observation {
	lines := strings.Split(instr, "\n")
	observations := make([]*Observation, len(lines))
	for idx, line := range lines {
		parts := strings.Split(line, "|")
		patterns := strings.Split(strings.Trim(parts[0], " "), " ")
		for i := 0; i < len(patterns); i += 1 {
			patterns[i] = SortString(patterns[i])
		}
		digits := strings.Split(strings.Trim(parts[1], " "), " ")
		for i := 0; i < len(digits); i += 1 {
			digits[i] = SortString(digits[i])
		}
		observations[idx] = &Observation{
			Patterns: patterns,
			Digits:   digits,
		}
	}
	return observations
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
