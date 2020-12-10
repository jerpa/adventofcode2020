package main

import (
	c "adventofcode2020/common"
	"sort"
)

func main() {

	data := c.GetInts(c.ReadInputFile())
	sort.Ints(data)

	diffOne := 0
	diffThree := 1

	last := 0
	for _, v := range data {
		if v-last == 1 {
			diffOne++
		} else if v-last == 3 {
			diffThree++
		}
		last = v
	}

	c.Print(diffOne, diffThree, diffOne*diffThree)

}
