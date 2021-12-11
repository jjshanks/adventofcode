package challenge

import "strings"

var closers = []rune{')', '}', ']', '>'}
var closerValuesPart1 = map[rune]int{')': 3, '}': 1197, ']': 57, '>': 25137}
var closerValuesPart2 = map[rune]int{')': 1, '}': 3, ']': 2, '>': 4}
var closerPairs = map[rune]rune{')': '(', '}': '{', ']': '[', '>': '<'}
var openerPairs = map[rune]rune{'(': ')', '{': '}', '[': ']', '<': '>'}

func parse(instr string) []string {
	return strings.Split(instr, "\n")
}

func contains(stack []rune, needle rune) bool {
	for _, i := range stack {
		if needle == i {
			return true
		}
	}
	return false
}
