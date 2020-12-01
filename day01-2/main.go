package main

import (
	c "adventofcode2020/common"
)

func main() {
	data := c.GetInts(c.ReadInputFile())

	for i := range data {
		for a := i + 1; a < len(data); a++ {
			for j := a + 1; j < len(data); j++ {
				if data[i]+data[a]+data[j] == 2020 {
					c.Print(data[i] * data[a] * data[j])
					return
				}
			}
		}
	}

}
