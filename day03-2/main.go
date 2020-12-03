package main

import (
	c "adventofcode2020/common"
)

type point struct {
	x int
	y int
}

func checkRoute(xStep, yStep int) int {
	data := c.ReadInputFile()
	d := 0
	x := 0

	for y := 0; y < len(data); y += yStep {
		if data[y][x%len(data[y])] == '#' {
			d++
		}
		x += xStep
	}
	return d
}

func main() {
	d := 1
	d *= checkRoute(1, 1)
	d *= checkRoute(3, 1)
	d *= checkRoute(5, 1)
	d *= checkRoute(7, 1)
	d *= checkRoute(1, 2)
	c.Print(d)

}
