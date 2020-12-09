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
	sum := 0

	preLength := 25
	data := c.GetInts(c.ReadInputFile())

	for prePos := 0; prePos < len(data)-preLength; prePos++ {
		preamble := data[prePos : prePos+preLength]
		if !validSum(preamble, data[prePos+preLength]) {
			c.Print(data[prePos+preLength])
			return
		}

	}
	c.Print(sum)

}
