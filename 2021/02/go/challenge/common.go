package challenge

import (
	"strconv"
	"strings"
)

type Instruction struct {
	Direction string
	Value     int
}

func parse(instr string) []Instruction {
	instructions := make([]Instruction, 0)
	for _, line := range strings.Split(instr, "\n") {
		parts := strings.Split(line, " ")
		value, _ := strconv.Atoi(parts[1])
		instructions = append(instructions, Instruction{
			Direction: parts[0],
			Value:     value,
		})
	}
	return instructions
}
