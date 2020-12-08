package main

import (
	c "adventofcode2020/common"
	"strings"
)

// test data and return value + true if it wasn't an infinite loop
func run(data []string) (int, bool) {
	sum := 0
	visited := map[int]bool{}
	pos := 0
	for {
		if pos >= len(data) {
			return sum, true
		}
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
	return sum, false
}

func main() {
	data := c.ReadInputFile()
	for i := range data {
		inf := strings.Split(data[i], " ")
		val := c.GetInt(inf[1])
		if inf[0] == "nop" && val != 0 {
			data[i] = strings.Replace(data[i], "nop", "jmp", 1)
			if v, t := run(data); t {
				c.Print(v)
				return
			}
			data[i] = strings.Replace(data[i], "jmp", "nop", 1)
		} else if inf[0] == "jmp" {
			data[i] = strings.Replace(data[i], "jmp", "nop", 1)
			if v, t := run(data); t {
				c.Print(v)
				return
			}
			data[i] = strings.Replace(data[i], "nop", "jmp", 1)
		}
	}

}
