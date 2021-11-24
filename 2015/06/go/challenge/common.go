package challenge

import (
	"regexp"
	"strconv"
	"strings"
)

/*
turn on 182,252 through 279,941
toggle 261,143 through 969,657
turn off 69,600 through 518,710
*/
func parse(instr string) []Instruction {
	r := regexp.MustCompile(`^([a-z ]+) ([\d]+),([\d]+) through ([\d]+),([\d]+)$`)
	parts := strings.Split(instr, "\n")
	instructions := make([]Instruction, 0, len(parts))
	for _, part := range parts {
		captures := r.FindStringSubmatch(part)
		codeStr := captures[1]
		var code Code
		switch codeStr {
		case "toggle":
			code = Toggle
		case "turn on":
			code = TurnOn
		case "turn off":
			code = Turnoff
		}
		startX, _ := strconv.Atoi(captures[2])
		startY, _ := strconv.Atoi(captures[3])
		endX, _ := strconv.Atoi(captures[4])
		endY, _ := strconv.Atoi(captures[5])
		start := Point{
			X: startX,
			Y: startY,
		}
		end := Point{
			X: endX,
			Y: endY,
		}
		instruction := Instruction{
			Code:  code,
			Start: start,
			End:   end,
		}
		instructions = append(instructions, instruction)
	}

	return instructions
}

type Point struct {
	X int
	Y int
}

type Instruction struct {
	Start Point
	End   Point
	Code  Code
}

type Code int

const (
	TurnOn Code = iota
	Turnoff
	Toggle
)

func NewGrid() [][]int {
	grid := make([][]int, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}
	return grid
}

func max(nums ...int) int {
	m := nums[0]
	for _, n := range nums {
		if n > m {
			m = n
		}
	}
	return m
}
