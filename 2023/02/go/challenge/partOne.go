package challenge

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	gameRegex, _ = regexp.Compile("^Game ([0-9]+):")
)

func PartOne(instr string) int {
	total := 0
	lines := strings.Split(instr, "\n")
	for _, line := range lines {
		validGame := true
		gameMatch := gameRegex.FindSubmatch([]byte(line))
		gameNum, _ := strconv.ParseInt(string(gameMatch[1]), 10, 32)
		draws := strings.Split(strings.Split(line, ":")[1], ";")
		for _, draw := range draws {
			cubes := strings.Split(draw, ",")
			for _, cube := range cubes {
				cube = strings.TrimSpace(cube)
				parts := strings.Split(cube, " ")
				color := parts[1]
				num, _ := strconv.ParseInt(parts[0], 10, 32)
				switch {
				case color == "red":
					if num > 12 {
						validGame = false
					}
				case color == "blue":
					if num > 14 {
						validGame = false
					}
				case color == "green":
					if num > 13 {
						validGame = false
					}
				}
			}
		}
		if validGame {
			total = total + int(gameNum)
		}
	}
	return total
}

//only 12 red cubes, 13 green cubes, and 14 blue cubes?
// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
// Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
// Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
