package main

import (
	c "adventofcode2020/common"
)

func main() {
	sum := 0

	for _, v := range c.GetInts(c.ReadInputFile()) {
		v++
	}
	c.Print(sum)

}
