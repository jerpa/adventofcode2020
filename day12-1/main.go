package main

import (
	c "adventofcode2020/common"
	"fmt"
)

func move(p *c.Point, curDir, instr string) string {
	ins := instr[0:1]
	ln := c.GetInt(instr[1:])
	if ins == "E" {
		p.X += ln
	} else if ins == "W" {
		p.X -= ln
	} else if ins == "N" {
		p.Y -= ln
	} else if ins == "S" {
		p.Y += ln
	} else if ins == "F" {
		move(p, curDir, fmt.Sprintf("%s%d", curDir, ln))
	} else if ins == "L" {
		for ln > 0 {
			if curDir == "N" {
				curDir = "W"
			} else if curDir == "W" {
				curDir = "S"
			} else if curDir == "S" {
				curDir = "E"
			} else if curDir == "E" {
				curDir = "N"
			}
			ln -= 90
		}
	} else if ins == "R" {
		for ln > 0 {
			if curDir == "N" {
				curDir = "E"
			} else if curDir == "W" {
				curDir = "N"
			} else if curDir == "S" {
				curDir = "W"
			} else if curDir == "E" {
				curDir = "S"
			}
			ln -= 90
		}
	}
	return curDir

}

func main() {
	p := c.Point{}
	dir := "E"

	for _, v := range c.ReadInputFile() {
		dir = move(&p, dir, v)
	}
	c.Print(p.Manhattan())

}
