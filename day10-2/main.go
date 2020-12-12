package main

import (
	c "adventofcode2020/common"
	"sort"
)

var checked map[[16]byte]uint64

func indexOf(data []int, num int) int {
	for i := range data {
		if i > 3 {
			return -1
		}
		if data[i] == num {
			return i
		}

	}
	return -1
}

func chain(data []int) uint64 {

	var res uint64

	if len(data) == 1 {
		return 1
	}
	if v, ok := checked[c.Checksum(data)]; ok {
		return v
	}

	if x := indexOf(data[1:], data[0]+1); x != -1 {
		v := chain(data[x+1:])
		checked[c.Checksum(data[x+1:])] = v
		res += v
	}
	if x := indexOf(data[1:], data[0]+2); x != -1 {
		v := chain(data[x+1:])
		checked[c.Checksum(data[x+1:])] = v
		res += v
	}
	if x := indexOf(data[1:], data[0]+3); x != -1 {
		v := chain(data[x+1:])
		checked[c.Checksum(data[x+1:])] = v
		res += v
	}

	return res
}

func main() {
	checked = map[[16]byte]uint64{}
	data := c.GetInts(c.ReadInputFile())
	sort.Ints(data)
	data = append([]int{0}, data...)
	sum := chain(data)

	c.Print(sum)

}
