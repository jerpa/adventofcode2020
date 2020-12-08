package main

import (
	c "adventofcode2020/common"
	"strings"
)

func main() {
	sum := 0
	data := c.ReadInputFile()
	visited := map[int]bool{}
	pos := 0
	for {
		if _, exists := visited[pos]; exists {
			break
		}
		visited[pos] = true
		inf := strings.Split(data[pos], " ")
		val := c.GetInt(inf[1])
		if inf[0] == "nop" {
			pos++
		} else if inf[0] == "acc" {
			sum += val
			pos++
		} else if inf[0] == "jmp" {
			pos += val
		}
	}
	c.Print(sum)

}
