package challenge

import (
	"container/heap"
)

func PartTwo(instr string) int {
	riskMap := expand(parse(instr))
	paths := &Paths{{X: 1, Y: 0, D: 0}, {X: 0, Y: 1, D: 0}}
	heap.Init(paths)
	cur := heap.Pop(paths).(Path)
	for !riskMap.IsEnd(cur.X, cur.Y) {
		n := riskMap.Neighbors(cur)
		for _, p := range n {
			heap.Push(paths, p)
		}
		cur = heap.Pop(paths).(Path)
	}
	return cur.D + int(riskMap.At(cur.X, cur.Y))
}
