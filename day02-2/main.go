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
		first, _ := strconv.Atoi(data[1])
		second, _ := strconv.Atoi(data[2])
		match := 0

		if string(data[4][first-1]) == data[3] {
			match++
		}
		if string(data[4][second-1]) == data[3] {
			match++
		}
		if match == 1 {
			d++
		}
	}
	c.Print(d)

}
