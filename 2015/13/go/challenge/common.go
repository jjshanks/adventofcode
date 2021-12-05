package challenge

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/dbyio/heappermutations"
)

var regex = regexp.MustCompile(`^([a-zA-Z]+) would (gain|lose) ([0-9]+) happiness units by sitting next to ([a-zA-Z]+)\.$`)

func fromLine(line string) (string, string, int) {
	matches := regex.FindStringSubmatch(line)
	delta, _ := strconv.Atoi(matches[3])
	if matches[2] == "lose" {
		delta = -delta
	}
	return matches[1], matches[4], delta
}

func parse(instr string) map[string]map[string]int {
	seatings := make(map[string]map[string]int)
	lines := strings.Split(instr, "\n")
	for _, line := range lines {
		left, right, delta := fromLine(line)

		leftMap, ok := seatings[left]
		if !ok {
			leftMap = make(map[string]int)
			seatings[left] = leftMap
		}
		leftMap[right] = delta
	}
	return seatings
}

func findMaxHappiness(seatings map[string]map[string]int, names []string) int {
	allArrangments := heappermutations.Strings(names)
	max := 0
	for _, arrangment := range allArrangments {
		total := 0
		for i, left := range arrangment {
			rightIdx := (i + 1) % len(arrangment)
			total += seatings[left][arrangment[rightIdx]]
			total += seatings[arrangment[rightIdx]][left]
		}
		if total > max {
			max = total
		}
	}
	return max
}
