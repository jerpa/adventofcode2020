package main

import (
	c "adventofcode2020/common"
)

func checkAdjacent(data []string, x, y int) int {
	sum := 0

	for yPos, xPos := y-1, x-1; yPos >= 0 && xPos >= 0; yPos-- {
		if data[yPos][xPos] == '#' {
			sum++
			break
		}
		if data[yPos][xPos] == 'L' {
			break
		}
		xPos--
	}
	for yPos, xPos := y-1, x; yPos >= 0 && xPos >= 0; yPos-- {
		if data[yPos][xPos] == '#' {
			sum++
			break
		}
		if data[yPos][xPos] == 'L' {
			break
		}
	}
	for yPos, xPos := y-1, x+1; yPos >= 0 && xPos < len(data[yPos]); yPos-- {
		if data[yPos][xPos] == '#' {
			sum++
			break
		}
		if data[yPos][xPos] == 'L' {
			break
		}
		xPos++
	}

	for yPos, xPos := y, x-1; xPos >= 0; xPos-- {
		if data[yPos][xPos] == '#' {
			sum++
			break
		}
		if data[yPos][xPos] == 'L' {
			break
		}
	}
	for yPos, xPos := y, x+1; xPos < len(data[yPos]); xPos++ {
		if data[yPos][xPos] == '#' {
			sum++
			break
		}
		if data[yPos][xPos] == 'L' {
			break
		}
	}

	for yPos, xPos := y+1, x-1; yPos < len(data) && xPos >= 0; yPos++ {
		if data[yPos][xPos] == '#' {
			sum++
			break
		}
		if data[yPos][xPos] == 'L' {
			break
		}
		xPos--
	}
	for yPos, xPos := y+1, x; yPos < len(data) && xPos >= 0; yPos++ {
		if data[yPos][xPos] == '#' {
			sum++
			break
		}
		if data[yPos][xPos] == 'L' {
			break
		}
	}
	for yPos, xPos := y+1, x+1; yPos < len(data) && xPos < len(data[yPos]); yPos++ {
		if data[yPos][xPos] == '#' {
			sum++
			break
		}
		if data[yPos][xPos] == 'L' {
			break
		}
		xPos++
	}

	return sum
}

func applyRules(data []string) []string {
	res := []string{}
	for y := range data {
		row := ""
		for x := 0; x < len(data[y]); x++ {
			if data[y][x] == '.' {
				row += "."
			} else if data[y][x] == 'L' {
				if checkAdjacent(data, x, y) == 0 {
					row += "#"
				} else {
					row += "L"
				}
			} else if data[y][x] == '#' {
				if checkAdjacent(data, x, y) >= 5 {
					row += "L"
				} else {
					row += "#"
				}
			}
		}
		res = append(res, row)
	}
	return res
}

func main() {

	data := c.ReadInputFile()

	for {
		res := applyRules(data)
		if c.Checksum(res) == c.Checksum(data) {
			break
		}
		data = res
	}

	c.Print(c.CountChar(data, '#'))

}
