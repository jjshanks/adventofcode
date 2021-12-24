package challenge

import (
	"container/list"
	"fmt"
	"regexp"
	"strings"
	"time"
)

type Key struct {
	L byte
	R byte
}

func parse(instr string) (map[Key]byte, []byte) {
	lines := strings.Split(instr, "\n")
	template := []byte(lines[0])
	lines = lines[2:]
	pairs := make(map[Key]byte)
	r := regexp.MustCompile(`([A-Z]+) -> ([A-Z])`)
	for _, line := range lines {
		captures := r.FindStringSubmatch(line)
		key := Key{
			L: captures[1][0],
			R: captures[1][1],
		}
		pairs[key] = captures[2][0]

	}
	return pairs, template
}

func newProcess(template []byte, pairs map[Key]byte, times int) int {
	if len(template) > 10 {
		return 2
	}
	p := make(map[byte]map[byte]byte)
	for k, v := range pairs {
		if _, ok := p[k.L]; !ok {
			p[k.L] = make(map[byte]byte)
		}
		p[k.L][k.R] = v
	}
	list := list.New()
	for _, b := range template {
		list.PushBack(b)
	}
	fmt.Println()
	for i := 0; i < times; i += 1 {
		fmt.Println(i, list.Len(), time.Now())
		cur := list.Front()
		for cur.Next() != nil {
			a := p[cur.Value.(byte)][cur.Next().Value.(byte)]
			// fmt.Print(cur.Value, ",", cur.Next().Value, " -> ", p[cur.Value.(byte)][cur.Next().Value.(byte)])
			// fmt.Println()
			cur = cur.Next()
			list.InsertBefore(a, cur)
		}
		// fmt.Println("----")
	}
	counts := make(map[byte]int)
	cur := list.Front()
	for cur.Next() != nil {
		counts[cur.Value.(byte)] += 1
	}
	most, least := 0, len(template)
	for _, v := range counts {
		if most < v {
			most = v
		}
		if least > v {
			least = v
		}
	}
	return most - least
}

func process(template []byte, pairs map[Key]byte, times int) int {
	for i := 0; i < times; i += 1 {
		newTemplate := make([]byte, 0, len(template)*2)
		for j := 0; j < len(template)-1; j += 1 {
			key := Key{
				L: template[j],
				R: template[j+1],
			}
			newTemplate = append(newTemplate, template[j])
			if v, ok := pairs[key]; ok {
				newTemplate = append(newTemplate, v)
			}
		}
		newTemplate = append(newTemplate, template[len(template)-1])
		template = newTemplate
	}
	counts := make(map[byte]int)
	for i := 0; i < len(template); i += 1 {
		counts[template[i]] += 1
	}
	most, least := 0, len(template)
	for _, v := range counts {
		if most < v {
			most = v
		}
		if least > v {
			least = v
		}
	}
	return most - least
}
