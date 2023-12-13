package challenge

import (
	"regexp"
	"strconv"
	"strings"
)

func PartOne(instr string) int {
	lines := strings.Split(instr, "\n")
	r := regexp.MustCompile(`([0-9])`)
	sum := 0
	for _, line := range lines {
		matches := r.FindAll([]byte(line), -1)
		first, _ := strconv.ParseInt(string(matches[0]), 10, 32)
		last, _ := strconv.ParseInt(string(matches[len(matches)-1]), 10, 32)
		sum = sum + int(first)*10 + int(last)
	}
	return sum
}
