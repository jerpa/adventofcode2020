package main

import (
	c "adventofcode2020/common"
	"math"
)

func getID(row, col int) int {
	return row*8 + col
}

func getVal(min, max int, pos string) (int, int) {
	mid := (float64(min) + float64(max)) / 2
	if pos == "F" || pos == "L" {
		// lower half
		return min, int(math.Floor(mid))
	} else if pos == "B" || pos == "R" {
		// upper half
		return int(math.Ceil(mid)), max
	}
	panic("invalid values")
}

func main() {
	data := c.ReadInputFile()
	maxID := 0
	places := map[int]bool{}

	for _, v := range data {
		rowMin := 0
		rowMax := 127
		colMin := 0
		colMax := 7
		for pos := 0; pos < 7; pos++ {
			rowMin, rowMax = getVal(rowMin, rowMax, string(v[pos]))
		}
		for pos := 7; pos < 10; pos++ {
			colMin, colMax = getVal(colMin, colMax, string(v[pos]))
		}
		id := getID(rowMin, colMin)
		places[id] = true
		if id > maxID {
			maxID = id
		}
	}
	for pos := 1; pos < maxID; pos++ {
		if _, ok := places[pos]; !ok {
			if _, before := places[pos-1]; !before {
				continue
			}
			if _, after := places[pos+1]; after {
				c.Print(pos)
			}
		}
	}

}
