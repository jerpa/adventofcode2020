package main

import (
	c "adventofcode2020/common"
)

func move(ship, waypoint *c.Point, instr string) {
	ins := instr[0:1]
	ln := c.GetInt(instr[1:])
	if ins == "E" {
		waypoint.X += ln
	} else if ins == "W" {
		waypoint.X -= ln
	} else if ins == "N" {
		waypoint.Y -= ln
	} else if ins == "S" {
		waypoint.Y += ln
	} else if ins == "F" {
		ship.X += ln * waypoint.X
		ship.Y += ln * waypoint.Y
	} else if ins == "L" {
		for ln > 0 {
			waypoint.X, waypoint.Y = waypoint.Y, -1*waypoint.X
			ln -= 90
		}
	} else if ins == "R" {
		for ln > 0 {
			waypoint.X, waypoint.Y = -1*waypoint.Y, waypoint.X
			ln -= 90
		}
	}
}

func main() {
	waypoint := c.Point{X: 10, Y: -1}
	ship := c.Point{}

	for _, v := range c.ReadInputFile() {
		move(&ship, &waypoint, v)
	}
	c.Print(ship.Manhattan())

}
