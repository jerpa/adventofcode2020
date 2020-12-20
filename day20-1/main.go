package main

import (
	c "adventofcode2020/common"
	"errors"
	"fmt"
	"strings"
)

type tile struct {
	data []string
	id   int
	x    int
	y    int
}

func coordID(x, y int) string {
	return fmt.Sprintf("%d:%d", x, y)
}

func iDcoord(id string) (int, int) {
	s := strings.Split(id, ":")
	return c.GetInt(s[0]), c.GetInt(s[1])
}

func (t *tile) rotate() {
	res := []string{}
	for x := 0; x < len(t.data[0]); x++ {
		row := ""
		for y := len(t.data) - 1; y >= 0; y-- {
			row += string(t.data[y][x])
		}
		res = append(res, row)
	}
	t.data = res
}
func (t *tile) flipX() {
	res := []string{}
	for _, y := range t.data {
		row := ""
		for _, x := range y {
			row = string(x) + row
		}
		res = append(res, row)
	}
	t.data = res
}

func find(tiles []tile, x, y int, edge string) (tile, error) {
	for t := range tiles {
		if tiles[t].x == 0 {
			for f := 0; f < 4; f++ {
				tiles[t].flipX()
				for i := 0; i < 4; i++ {
					if tiles[t].data[0] == edge {
						tiles[t].x = x
						tiles[t].y = y
						return tiles[t], nil
					}
					tiles[t].rotate()
				}
			}
		}
	}
	return tile{}, errors.New("no match")
}

func main() {
	sum := 0
	t := tile{}
	m := []tile{}
	res := map[string]tile{}
	for _, v := range c.ReadInputFile() {
		if len(v) == 0 {
			m = append(m, t)
			t = tile{}
		} else if v[0] == 'T' {
			v = strings.TrimLeft(v, "Tile ")
			v = strings.Trim(v, ":")
			t.id = c.GetInt(v)
		} else {
			t.data = append(t.data, v)
		}
	}
	m = append(m, t)
	m[0].x = 1000
	m[0].y = 1000
	res[coordID(1000, 1000)] = m[0]
	for len(res) < len(m) {
		for k, v := range res {
			x, y := iDcoord(k)
			if _, ok := res[coordID(x+1, y)]; !ok {
				s := ""
				for _, c := range v.data {
					s += string(c[len(c)-1])
				}
				n, err := find(m, x+1, y, s)
				if err == nil {
					n.rotate()
					n.flipX()
					res[coordID(x+1, y)] = n
				}
			}
			if _, ok := res[coordID(x-1, y)]; !ok {
				s := ""
				for _, c := range v.data {
					s += string(c[0])
				}
				n, err := find(m, x-1, y, s)
				if err == nil {
					n.rotate()
					res[coordID(x-1, y)] = n
				}
			}
			if _, ok := res[coordID(x, y-1)]; !ok {
				n, err := find(m, x, y-1, v.data[0])
				if err == nil {
					n.flipX()
					n.rotate()
					n.rotate()
					res[coordID(x, y-1)] = n
				}
			}
			if _, ok := res[coordID(x, y+1)]; !ok {
				n, err := find(m, x, y+1, v.data[len(v.data)-1])
				if err == nil {
					res[coordID(x, y+1)] = n
				}
			}

		}

	}
	minX, minY, maxX, maxY := 0, 0, 0, 0
	for k := range res {
		x, y := iDcoord(k)
		if minX == 0 || x < minX {
			minX = x
		}
		if maxX == 0 || x > maxX {
			maxX = x
		}
		if minY == 0 || y < minY {
			minY = y
		}
		if maxY == 0 || y > maxY {
			maxY = y
		}
	}
	sum = res[coordID(minX, minY)].id * res[coordID(maxX, minY)].id * res[coordID(maxX, maxY)].id * res[coordID(minX, maxY)].id

	c.Print(sum)

}
