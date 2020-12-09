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

func sum(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func minInt(nums []int) int {
	res := 0
	for i, v := range nums {
		if i == 0 || v < res {
			res = v
		}
	}
	return res
}
func maxInt(nums []int) int {
	res := 0
	for i, v := range nums {
		if i == 0 || v > res {
			res = v
		}
	}
	return res
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
			s := sum(data[start:stop])
			if s > inv {
				break
			}
			if s == inv {
				c.Print(minInt(data[start:stop]) + maxInt(data[start:stop]))
				return
			}
		}
		start++
	}

}
