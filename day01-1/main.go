package main

import (
	c "adventofcode2020/common"
)

func main() {
	data := c.GetInts(c.ReadInputFile())

	for i := range data {
		for a := i + 1; a < len(data); a++ {
			if data[i]+data[a] == 2020 {
				c.Print(data[i] * data[a])
				return
			}
		}
	}

}
