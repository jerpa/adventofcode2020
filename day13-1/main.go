package main

import (
	c "adventofcode2020/common"
	"strings"
)

func main() {
	val := 0
	data := c.ReadInputFile()
	t := c.GetInt(data[0])

	for _, v := range strings.Split(data[1], ",") {
		if v == "x" {
			continue
		}
		id := c.GetInt(v)
		//c.Print(id, t%id, id-(t%id))
		if val == 0 || id-(t%id) < val-(t%val) {
			val = id
		}
	}
	c.Print(val, val*(val-(t%val)))

}
