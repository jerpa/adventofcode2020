package main

import (
	c "adventofcode2020/common"
)

func checkAdjacent(data []string, x, y int) int {
	sum := 0

	for yPos := -1; yPos <= 1; yPos++ {
		for xPos := -1; xPos <= 1; xPos++ {
			if yPos == 0 && xPos == 0 {
				continue
			}
			if y+yPos < 0 || y+yPos >= len(data) {
				continue
			}
			if x+xPos < 0 || x+xPos >= len(data[y]) {
				continue
			}
			if data[y+yPos][x+xPos] == '#' {
				sum++
			}
		}
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
				if checkAdjacent(data, x, y) >= 4 {
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
