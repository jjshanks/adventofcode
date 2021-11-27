package challenge

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func parse(instr string) []Instruction {
	lines := strings.Split(instr, "\n")
	r := regexp.MustCompile(`^([a-zA-Z0-9 ]+) -> ([a-z]+)$`)
	instructions := make([]Instruction, 0, len(lines))
	for _, line := range lines {
		captures := r.FindStringSubmatch(line)
		left := captures[1]
		right := captures[2]
		instruction := Instruction{
			Out: right,
		}
		leftParts := strings.Split(left, " ")
		wirePattern := regexp.MustCompile("^[a-z]+$")
		switch len(leftParts) {
		case 1:
			instruction.Code = IDENTITY
			isWire := wirePattern.MatchString(leftParts[0])
			value, _ := strconv.Atoi(leftParts[0])
			instruction.In1 = Input{
				IsWire:   isWire,
				WireName: leftParts[0],
				Value:    uint16(value),
			}
		case 2:
			instruction.Code = Code(leftParts[0])
			isWire := wirePattern.MatchString(leftParts[1])
			value, _ := strconv.Atoi(leftParts[1])
			instruction.In1 = Input{
				IsWire:   isWire,
				WireName: leftParts[1],
				Value:    uint16(value),
			}
		case 3:
			isWire := wirePattern.MatchString(leftParts[0])
			value, _ := strconv.Atoi(leftParts[0])
			instruction.In1 = Input{
				IsWire:   isWire,
				WireName: leftParts[0],
				Value:    uint16(value),
			}
			instruction.Code = Code(leftParts[1])
			isWire2 := wirePattern.MatchString(leftParts[2])
			value2, _ := strconv.Atoi(leftParts[2])
			instruction.In2 = Input{
				IsWire:   isWire2,
				WireName: leftParts[2],
				Value:    uint16(value2),
			}
		}
		instruction.ID = uuid.New().String()
		instructions = append(instructions, instruction)
	}
	return instructions
}

type Instruction struct {
	In1  Input
	In2  Input
	Code Code
	Out  string
	ID   string
}

func (i Instruction) eval(wires map[string]uint16) (uint16, error) {
	val1 := uint16(0)
	if i.In1.IsWire {
		if _, ok := wires[i.In1.WireName]; !ok {
			return 0, errors.New("no wire value yet")
		}
		val1 = wires[i.In1.WireName]
	} else {
		val1 = i.In1.Value
	}
	val2 := uint16(0)
	_, val2ok := wires[i.In2.WireName]
	if i.In2.IsWire {
		val2 = wires[i.In2.WireName]
	} else {
		val2 = i.In2.Value
	}

	switch i.Code {
	case IDENTITY:
		return val1, nil
	case NOT:
		return ^val1, nil
	case LSHIFT:
		return val1 << i.In2.Value, nil
	case RSHIFT:
		return val1 >> i.In2.Value, nil
	case AND:
		if !val2ok {
			return 0, errors.New("no wire value yet")
		}
		return val1 & val2, nil
	case OR:
		if !val2ok {
			return 0, errors.New("no wire value yet")
		}
		return val1 | val2, nil
	}
	return 0, nil
}

type Input struct {
	IsWire   bool
	WireName string
	Value    uint16
}

type Code string

const (
	AND      Code = "AND"
	OR       Code = "OR"
	LSHIFT   Code = "LSHIFT"
	RSHIFT   Code = "RSHIFT"
	NOT      Code = "NOT"
	IDENTITY Code = "IDENTITY"
)
