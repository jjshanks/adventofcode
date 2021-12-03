package challenge

import (
	"encoding/json"
	"fmt"
)

func parse(instr string) interface{} {
	b := []byte(instr)
	var f interface{}
	json.Unmarshal(b, &f)
	return f
}

func processNode(node interface{}, ignore string) int {
	total := 0
	switch t := node.(type) {
	case []interface{}:
		for _, v := range t {
			total += processNode(v, ignore)
		}
	case map[string]interface{}:
		for _, v := range t {
			if v == ignore {
				return 0
			}
			total += processNode(v, ignore)
		}
	case float64:
		return int(t)
	case string:
		return 0
	default:
		fmt.Printf("unkown -- %v\n", t)
	}
	return total
}
