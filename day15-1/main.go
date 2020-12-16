package main

import (
	c "adventofcode2020/common"
	"strings"
)

func main() {

	data := c.GetInts(strings.Split(c.ReadInputFile()[0], ","))
mainLoop:
	for i := len(data); i < 2020; i++ {
		last := data[i-1]
		for a := i - 2; a >= 0; a-- {
			if data[a] == last {
				data = append(data, i-a-1)
				continue mainLoop
			}
		}
		data = append(data, 0)
	}
	c.Print(data[2019])

}
