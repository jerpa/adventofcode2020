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
func getID(x, y int, dir string) string {
	if dir == "w" {
		x--
	} else if dir == "e" {
		x++
	} else if dir == "se" {
		if y%2 == 1 || y%2 == -1 {
			x++
		}
		y--

	} else if dir == "sw" {
		if y%2 == 0 {
			x--
		}
		y--
	} else if dir == "ne" {
		if y%2 == 1 || y%2 == -1 {
			x++
		}
		y++
	} else if dir == "nw" {
		if y%2 == 0 {
			x--
		}
		y++
	}
	return coordID(x, y)
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
	s := 0
	for _, v := range tiles {
		if v {
			s++
		}
	}
	c.Print(s)
	adjOk := map[string]bool{}
	for i := 0; i < 100; i++ {
		for k := range tiles {
			// Don't check tiles that we already know has neighbors
			if adjOk[k] {
				continue
			}
			// Fill the missing neighbors with white
			x, y := iDcoord(k)
			if _, ok := tiles[getID(x, y, "nw")]; !ok {
				tiles[getID(x, y, "nw")] = false
			}
			if _, ok := tiles[getID(x, y, "ne")]; !ok {
				tiles[getID(x, y, "ne")] = false
			}
			if _, ok := tiles[getID(x, y, "e")]; !ok {
				tiles[getID(x, y, "e")] = false
			}
			if _, ok := tiles[getID(x, y, "se")]; !ok {
				tiles[getID(x, y, "se")] = false
			}
			if _, ok := tiles[getID(x, y, "sw")]; !ok {
				tiles[getID(x, y, "sw")] = false
			}
			if _, ok := tiles[getID(x, y, "w")]; !ok {
				tiles[getID(x, y, "w")] = false
			}
			adjOk[k] = true
		}
		copy := map[string]bool{}
		for k, v := range tiles {
			copy[k] = v
		}
		for k, v := range copy {
			x, y := iDcoord(k)
			n := 0
			if m, ok := copy[getID(x, y, "nw")]; ok && m {
				n++
			}
			if m, ok := copy[getID(x, y, "ne")]; ok && m {
				n++
			}
			if m, ok := copy[getID(x, y, "e")]; ok && m {
				n++
			}
			if m, ok := copy[getID(x, y, "se")]; ok && m {
				n++
			}
			if m, ok := copy[getID(x, y, "sw")]; ok && m {
				n++
			}
			if m, ok := copy[getID(x, y, "w")]; ok && m {
				n++
			}
			if v && (n == 0 || n > 2) {
				tiles[k] = false
			} else if !v && n == 2 {
				tiles[k] = true
			}
		}
		s := 0
		for _, v := range tiles {
			if v {
				s++
			}
		}
		c.Print(i+1, s)
	}

	for _, v := range tiles {
		if v {
			sum++
		}
	}
	c.Print(sum)

}
