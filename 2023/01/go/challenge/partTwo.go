package challenge

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	valueMap = map[string]int64{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
)

func PartTwo(instr string) int {
	lines := strings.Split(instr, "\n")
	r := regexp.MustCompile(`([0-9]|one|two|three|four|five|six|seven|eight|nine)`)
	sum := 0
	for _, line := range lines {
		matches := r.FindAll([]byte(line), -1)
		firstStr := string(matches[0])
		var first int64
		if _, ok := valueMap[firstStr]; ok {
			first = valueMap[firstStr]
		} else {
			first, _ = strconv.ParseInt(firstStr, 10, 32)
		}
		lastStr := string(matches[len(matches)-1])
		var last int64
		if _, ok := valueMap[lastStr]; ok {
			last = valueMap[lastStr]
		} else {
			last, _ = strconv.ParseInt(lastStr, 10, 32)
		}
		sum = sum + int(first)*10 + int(last)
	}
	return sum
}
