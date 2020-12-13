package main

import (
	c "adventofcode2020/common"
	"strings"
)

func main() {

	data := c.ReadInputFile()

	busOffset := map[int]int{}
	bus := []int{}

	for i, v := range strings.Split(data[1], ",") {
		if v == "x" {
			continue
		}
		busOffset[c.GetInt(v)] = i
		bus = append(bus, c.GetInt(v))
	}
	step := bus[0]
	t := 0
	for {
		step = bus[0]
		for i := 1; i < len(bus); i++ {
			id := bus[i]

			if ((t + busOffset[id]) % id) != 0 {
				break
			}
			step *= id
			if i == len(bus)-1 {
				// Match!
				c.Print(t)
				return
			}
		}
		t += step
	}
}
