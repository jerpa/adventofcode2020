package main

import (
	c "adventofcode2020/common"
)

func countMatch(data map[rune]int, num int) int {
	res := 0
	for _, v := range data {
		if v == num {
			res++
		}
	}
	return res
}

func main() {
	sum := 0

	selected := map[rune]int{}
	cnt := 0
	for _, v := range c.ReadInputFile() {
		if len(v) == 0 {
			sum += countMatch(selected, cnt)
			selected = map[rune]int{}
			cnt = 0
		} else {
			cnt++
			for _, b := range v {
				selected[b]++
			}
		}
	}
	sum += countMatch(selected, cnt)

	c.Print(sum)

}
