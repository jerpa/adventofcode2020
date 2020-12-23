package main

import (
	c "adventofcode2020/common"
)

func indexOf(sl []int, val int) int {
	for i := range sl {
		if sl[i] == val {
			return i
		}
	}
	return -1
}

func main() {
	sum := 0
	data := "284573961"
	nums := []int{}
	curr := 3

	for _, v := range data {
		nums = append(nums, c.GetInt(string(v)))
	}
	for i := 0; i < 100; i++ {
		rem := nums[1:4]
		c.Print(nums)
		c.Print(rem)
		nums = append([]int{nums[0]}, nums[4:]...)
		val := nums[0] - 1
		if val == 0 {
			val = 9
		}
		for indexOf(rem, val) != -1 {
			val--
			if val == 0 {
				val = 9
			}
		}
		pos := indexOf(nums, val)
		c.Print(val, pos)
		nums = append(nums[:pos+1], append(rem, nums[pos+1:]...)...)
		curr = 1
		//c.Print(nums)

		nums = append(nums[curr:], nums[:curr]...)
	}
	c.Print(nums)
	pos := indexOf(nums, 1)
	nums = append(nums[pos:], nums[:pos]...)
	c.Print(nums[1:])
	c.Print(sum)

}
