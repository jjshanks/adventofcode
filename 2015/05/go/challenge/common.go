package challenge

import (
	"regexp"
	"strings"
)

func parse(instr string) []string {
	return []string{}
}

/*
It contains a pair of any two letters that appears at least twice in the string without overlapping, like xyxy (xy) or aabcdefgaa (aa), but not like aaa (aa, but it overlaps).
It contains at least one letter which repeats with exactly one letter between them, like xyx, abcdefeghi (efe), or even aaa.
*/
func IsNicePartTwo(s string) bool {
	doubleDouble := false
	for i, _ := range s {
		if doubleDouble || i+2 >= len(s) {
			break
		}
		target := s[i : i+2]
		ss := s[i+2:]
		doubleDouble = strings.Contains(ss, target)
	}
	doubleGap := false
	for i, c := range []byte(s) {
		if doubleGap || i+2 == len(s) {
			break
		}
		doubleGap = c == s[i+2]

	}
	return doubleDouble && doubleGap
}

/*
It contains at least three vowels (aeiou only), like aei, xazegov, or aeiouaeiouaeiou.
It contains at least one letter that appears twice in a row, like xx, abcdde (dd), or aabbccdd (aa, bb, cc, or dd).
It does not contain the strings ab, cd, pq, or xy, even if they are part of one of the other requirements.
*/
func IsNice(s string) bool {
	vowels := strings.Count(s, "a") + strings.Count(s, "e") + strings.Count(s, "i") + strings.Count(s, "o") + strings.Count(s, "u")
	doubleLetter := false
	for i, c := range []byte(s) {
		if doubleLetter || i+1 == len(s) {
			break
		}
		doubleLetter = c == s[i+1]

	}
	doesNotContain := regexp.MustCompile(`(ab|cd|pq|xy)`)
	return vowels >= 3 && doubleLetter && !doesNotContain.MatchString(s)
}
