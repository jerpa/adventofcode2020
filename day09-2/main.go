package main

import (
	c "adventofcode2020/common"
)

func validSum(num []int, sum int) bool {
	for i := range num {
		for u := i + 1; u < len(num); u++ {
			if num[i]+num[u] == sum {
				return true
			}
		}
	}
	return false
}

func main() {
	var inv int

	preLength := 25
	data := c.GetInts(c.ReadInputFile())

	for prePos := 0; prePos < len(data)-preLength; prePos++ {
		preamble := data[prePos : prePos+preLength]
		if !validSum(preamble, data[prePos+preLength]) {
			c.Print(data[prePos+preLength])
			inv = data[prePos+preLength]
			break
		}
	}

	start := 0
	stop := 0
	for {
		for stop = start + 1; stop < len(data); stop++ {
			s := c.Sum(data[start:stop])
			if s > inv {
				break
			}
			if s == inv {
				c.Print(c.MinInt(data[start:stop]) + c.MaxInt(data[start:stop]))
				return
			}
		}
		start++
	}

}
