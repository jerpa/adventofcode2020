package main

import (
	c "adventofcode2020/common"
)

// linked list
type item struct {
	val  int
	next *item
}

func main() {
	var start *item
	var last *item
	sum := 0
	data := "284573961"
	index := map[int]*item{}

	// fill with input
	for _, v := range data {
		x := item{val: c.GetInt(string(v))}
		index[x.val] = &x
		if start == nil {
			start = &x
		}
		if last != nil {
			last.next = &x
		}
		last = &x
	}
	// fill up to 1000000
	for i := 10; i <= 1000000; i++ {
		x := item{val: i}
		index[x.val] = &x
		last.next = &x

		last = &x

	}
	// connect end to start
	last.next = start

	curr := start

	var pos *item
	var rem *item
	var val int
	var x int

	for i := 0; i < 10000000; i++ {

		// pos=next after removed
		pos = curr.next.next.next.next

		// rem=start of removed
		rem = curr.next

		// connect curent to the one after removed
		curr.next = pos
		// find value to look for
		val = curr.val - 1
		if val == 0 {
			val = 1000000
		}
		// check so value isn't removed
		for x = 0; x < 3; x++ {
			if rem.val == val || rem.next.val == val || rem.next.next.val == val {
				val--
				if val == 0 {
					val = 1000000
				}
			} else {
				break
			}
		}
		// lookup in index for position of searched value
		pos = index[val]

		// insert the removed after newly found position
		rem.next.next.next = pos.next
		pos.next = rem
		// move to the next in line
		curr = curr.next
	}
	// where is number one?
	start = index[1]

	// find the one after and the one after that
	c.Print(start.next.val, start.next.next.val)
	sum = start.next.val * start.next.next.val
	c.Print(sum)

}
