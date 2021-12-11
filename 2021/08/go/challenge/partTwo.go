package challenge

import (
	"strings"
)

func PartTwo(instr string) int {
	inputSlice := parse(instr)
	sum := 0
	for _, o := range inputSlice {
		one, four, seven, eight := "", "", "", ""
		for _, p := range o.Patterns {
			if len(p) == 2 {
				one = p
			}
			if len(p) == 3 {
				seven = p
			}
			if len(p) == 4 {
				four = p
			}
			if len(p) == 7 {
				eight = p
			}
		}
		three, nine := "", ""
		// extra from one and sever is top line
		//
		for _, p := range o.Patterns {
			if len(p) == 6 {
				r := diff(p, one, four, seven)
				if len(r) == 1 {
					nine = p
				}
			}
			if len(p) == 5 {
				r := diff(p, one)
				if len(r) == 3 {
					three = p
				}
			}
		}
		zero, two, five, six := "", "", "", ""
		for _, p := range o.Patterns {
			if len(p) == 5 {
				if p == three {
					continue
				}
				// two five
				r := diff(p, nine)
				if len(r) == 0 {
					five = p
				}
				if len(r) == 1 {
					two = p
				}
			}
			if len(p) == 6 {
				if p == nine {
					continue
				}
				r := diff(p, one)
				if len(r) == 4 {
					zero = p
				}
				if len(r) == 5 {
					six = p
				}
			}
		}
		m := make(map[string]int)
		m[zero] = 0
		m[one] = 1
		m[two] = 2
		m[three] = 3
		m[four] = 4
		m[five] = 5
		m[six] = 6
		m[seven] = 7
		m[eight] = 8
		m[nine] = 9
		total := 0
		for _, d := range o.Digits {
			total *= 10
			total += m[d]
		}
		sum += total
	}
	return sum
}

func diff(main string, diffs ...string) string {
	r := ""
	for _, c := range strings.Split(main, "") {
		found := false
		for _, d := range diffs {
			if strings.Contains(d, c) {
				found = true
				break
			}
		}
		if !found {
			r += c
		}
	}
	return r
}

func detectTop(one, seven string) string {
	r := ""
	for _, c := range strings.Split(seven, "") {
		if !strings.Contains(one, c) {
			r = c
			break
		}
	}
	return r
}

func detectUpperLeftAndMiddle(one, four string) []string {
	r := make([]string, 0, 2)
	for _, c := range strings.Split(four, "") {
		if !strings.Contains(one, c) {
			r = append(r, c)
		}
	}
	return r
}
