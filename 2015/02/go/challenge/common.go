package challenge

import (
	"strconv"
	"strings"
)

type Box struct {
	Length int
	Width  int
	Height int
}

func (b *Box) PaperNeeded() int {
	return 2*b.Length*b.Width + 2*b.Width*b.Height + 2*b.Height*b.Length + min(b.Length*b.Width, b.Length*b.Height, b.Width*b.Height)
}

func (b *Box) RibbonNeeded() int {
	return 2*min(b.Length+b.Width, b.Length+b.Height, b.Width+b.Height) + b.Height*b.Length*b.Width
}

func min(nums ...int) int {
	m := nums[0]
	for _, n := range nums {
		if n < m {
			m = n
		}
	}
	return m
}

func parse(instr string) []*Box {
	boxStr := strings.Split(instr, "\n")
	boxes := make([]*Box, 0, len(boxStr))
	for _, bs := range boxStr {
		boxPartStr := strings.Split(bs, "x")
		l, _ := strconv.Atoi(boxPartStr[0])
		w, _ := strconv.Atoi(boxPartStr[1])
		h, _ := strconv.Atoi(boxPartStr[2])
		boxes = append(boxes, &Box{
			Length: l,
			Width:  w,
			Height: h,
		})
	}
	return boxes
}
