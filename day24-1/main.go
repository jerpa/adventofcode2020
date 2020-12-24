package main

import (
	c "adventofcode2020/common"
	"fmt"
	"strings"
)

/*

   /   \ /   \ /   \ /   \
  |-2,1 |-1,1 | 0,1 | 1,1 |
   \   / \   / \   / \   /
     |-1,0 | 0,0 | 1,0 |
   /   \ /   \ /   \ /   \
  |-2,-1|-1,-1| 0,-1| 1,-1|
   \   / \   / \   / \   /

   W=X-1
   E=X+1
   NW=X-1 (if Y%2==0), Y+1
   NE=X+1 (if Y%2==1 or -1), Y+1
   SW=X-1 (if Y%2==0), Y-1
   SE=X+1 (if Y%2==1 or -1), Y-1

*/

func coordID(x, y int) string {
	return fmt.Sprintf("%d:%d", x, y)
}

func iDcoord(id string) (int, int) {
	s := strings.Split(id, ":")
	return c.GetInt(s[0]), c.GetInt(s[1])
}

func main() {
	sum := 0
	tiles := map[string]bool{}

	for _, v := range c.ReadInputFile() {
		x, y := 0, 0

		for pos := 0; pos < len(v); pos++ {
			if v[pos] == 'w' {
				x--
			} else if v[pos] == 'e' {
				x++
			} else if v[pos] == 's' && v[pos+1] == 'e' {
				if y%2 == 1 || y%2 == -1 {
					x++
				}
				y--
				pos++
			} else if v[pos] == 's' && v[pos+1] == 'w' {
				if y%2 == 0 {
					x--
				}
				y--
				pos++
			} else if v[pos] == 'n' && v[pos+1] == 'e' {
				if y%2 == 1 || y%2 == -1 {
					x++
				}
				y++
				pos++
			} else if v[pos] == 'n' && v[pos+1] == 'w' {
				if y%2 == 0 {
					x--
				}
				y++
				pos++
			}
		}
		if _, ok := tiles[coordID(x, y)]; !ok {
			tiles[coordID(x, y)] = false
		}
		tiles[coordID(x, y)] = !tiles[coordID(x, y)]
	}
	for _, v := range tiles {
		if v {
			sum++
		}
	}
	c.Print(sum)

}
