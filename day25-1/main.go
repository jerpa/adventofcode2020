package main

import (
	c "adventofcode2020/common"
)

func getLoopSize(key int) int {
	val := 1

	for num := 1; ; num++ {
		val *= 7
		val %= 20201227
		if val == key {
			return num
		}
	}
}

func getEncryptionKey(loopSize, value int) int {
	val := 1
	for i := 0; i < loopSize; i++ {
		val *= value
		val %= 20201227
	}
	return val
}

func main() {
	sum := 0

	data := c.GetInts(c.ReadInputFile())
	cardLoop := getLoopSize(data[0])
	doorLoop := getLoopSize(data[1])
	c.Print(cardLoop, doorLoop)
	c.Print(getEncryptionKey(cardLoop, data[1]), getEncryptionKey(doorLoop, data[0]))

	for _, v := range c.GetInts(c.ReadInputFile()) {
		v++
	}
	c.Print(sum)

}
