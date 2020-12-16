package main

import (
	c "adventofcode2020/common"
	"strings"
)

func main() {
	prev := map[int]int{}
	spoken := map[int]int{}
	count := map[int]int{}
	var last int
	data := c.GetInts(strings.Split(c.ReadInputFile()[0], ","))

	for i, v := range data {
		spoken[v] = i
		prev[v] = i
		count[v] = 1
		last = v
	}

	for i := len(data); i < 30000000; i++ {
		if count[last] == 1 {
			last = 0
		} else {
			last = spoken[last] - prev[last]
		}
		prev[last] = spoken[last]
		spoken[last] = i
		count[last]++
	}

	c.Print(last)

}
