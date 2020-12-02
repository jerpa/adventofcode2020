package main

import (
	c "adventofcode2020/common"
	"regexp"
	"strconv"
)

func main() {
	s := c.ReadInputFile()
	d := 0
	reg := regexp.MustCompile(`(\d+)-(\d+)\s([a-z]):\s([a-z]+)`)
	for _, v := range s {
		data := reg.FindStringSubmatch(v)

		if data == nil {
			panic("no match")
		}
		min, _ := strconv.Atoi(data[1])
		max, _ := strconv.Atoi(data[2])
		match := 0
		for _, c := range data[4] {
			if string(c) == data[3] {
				match++
			}
		}
		if min <= match && match <= max {
			d++
		}

	}
	c.Print(d)

}
