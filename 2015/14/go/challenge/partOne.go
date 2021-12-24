package challenge

import "fmt"

func PartOne(instr string) int {
	reindeers := parse(instr)
	distances := make(map[string]int)
	for t := 1; t < 1000; t += 1 {
		for _, r := range reindeers {
			speed := r.SpeedAtTime(t)
			distances[r.Name] += speed
			if speed > 0 {
				fmt.Printf("%s has gone %d in %d\n", r.Name, distances[r.Name], t)
			}

		}
	}
	fmt.Printf("\n%v\n", distances)
	return 0
}
