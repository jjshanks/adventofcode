package challenge

import (
	"fmt"
	"strconv"
	"strings"
)

type Reindeer struct {
	Name     string
	Speed    int
	RunTime  int
	RestTime int
}

func (r *Reindeer) String() string {
	return fmt.Sprintf("%s, %d, %d, %d", r.Name, r.Speed, r.RunTime, r.RestTime)
}

func (r *Reindeer) SpeedAtTime(time int) int {
	window := time % (r.RunTime + r.RestTime + 1)
	if window > r.RunTime {
		return 0
	}
	return r.Speed
}

func parse(instr string) []*Reindeer {
	lines := strings.Split(instr, "\n")
	reindeers := make([]*Reindeer, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, " ")
		name := parts[0]
		speed, _ := strconv.Atoi(parts[3])
		runTime, _ := strconv.Atoi(parts[6])
		restTime, _ := strconv.Atoi(parts[13])
		reindeers = append(reindeers, &Reindeer{
			Name:     name,
			Speed:    speed,
			RunTime:  runTime,
			RestTime: restTime,
		})
	}
	return reindeers
}
