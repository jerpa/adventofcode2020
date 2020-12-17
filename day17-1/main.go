package main

import (
	c "adventofcode2020/common"
	"fmt"
	"strings"
)

func coordID(x, y, z int) string {
	return fmt.Sprintf("%d:%d:%d", x, y, z)
}

func iDcoord(id string) (int, int, int) {
	s := strings.Split(id, ":")
	return c.GetInt(s[0]), c.GetInt(s[1]), c.GetInt(s[2])
}

func countNeighbours(cube map[string]bool, posX, posY, posZ int) int {

	num := 0
	for z := -1; z <= 1; z++ {
		for y := -1; y <= 1; y++ {
			for x := -1; x <= 1; x++ {
				if z == 0 && y == 0 && x == 0 {
					continue
				}
				if cube[coordID(posX+x, posY+y, posZ+z)] {
					num++
				}
			}
		}
	}
	return num
}

func main() {
	sum := 0

	cube := map[string]bool{}
	for y, v := range c.ReadInputFile() {
		for x, xV := range v {
			if xV == '#' {
				cube[coordID(x, y, 0)] = true
			}
		}
	}

	for i := 0; i < 6; i++ {
		nC := map[string]bool{}
		for k := range cube {
			posX, posY, posZ := iDcoord(k)
			for z := -1; z <= 1; z++ {
				for y := -1; y <= 1; y++ {
					for x := -1; x <= 1; x++ {
						if z == 0 && y == 0 && x == 0 {
							continue
						}
						num := countNeighbours(cube, posX+x, posY+y, posZ+z)
						if cube[coordID(posX+x, posY+y, posZ+z)] && (num == 2 || num == 3) {
							nC[coordID(posX+x, posY+y, posZ+z)] = true
						} else if !cube[coordID(posX+x, posY+y, posZ+z)] && num == 3 {
							nC[coordID(posX+x, posY+y, posZ+z)] = true
						}
					}
				}
			}

		}
		cube = nC
	}

	for _, v := range cube {
		if v {
			sum++
		}
	}
	c.Print(sum)

}
