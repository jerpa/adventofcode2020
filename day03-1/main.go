package main

import (
	c "adventofcode2020/common"
)

type point struct {
	x int
	y int
}

func main() {
	data := c.ReadInputFile()
	d := 0
	x := 0
	for _, v := range data {
		if v[x%len(v)] == '#' {
			d++
		}
		x += 3
	}
	c.Print(d)

}
