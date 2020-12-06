package main

import (
	c "adventofcode2020/common"
)

func main() {
	sum := 0

	selected := map[rune]bool{}

	for _, v := range c.ReadInputFile() {
		if len(v) == 0 {
			sum += len(selected)
			selected = map[rune]bool{}
		} else {
			for _, b := range v {
				selected[b] = true
			}
		}
	}
	sum += len(selected)

	c.Print(sum)

}
