package main

import (
	c "adventofcode2020/common"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var sum int64
	mem := map[int64]int64{}
	mask := ""
	data := c.ReadInputFile()

	reg := regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)
	for i := range data {
		if v := strings.Split(data[i], " = "); v[0] == "mask" {
			mask = v[1]
		} else {

			m := reg.FindStringSubmatch(data[i])
			addr, _ := strconv.ParseInt(m[1], 10, 64)
			val, _ := strconv.ParseInt(m[2], 10, 64)
			var maskVal int64 = 1
			for mPos := len(mask) - 1; mPos >= 0; mPos-- {
				if mask[mPos] == '1' {
					val = val | maskVal
				} else if mask[mPos] == '0' && val&maskVal > 0 {
					val -= maskVal
				}
				maskVal *= 2
			}
			mem[addr] = val
		}
	}

	for _, v := range mem {

		sum += v
	}

	c.Print(sum)

}
