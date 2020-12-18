package main

import (
	c "adventofcode2020/common"
	"fmt"
	"strings"
)

func calc(s string) int {
	s = " " + s + " "
	p := strings.Index(s, "(")
	if p != -1 {
		pC := 1
		var e int
		for e = p + 1; e < len(s) && pC > 0; e++ {
			if s[e] == '(' {
				pC++
			}
			if s[e] == ')' {
				pC--
			}
		}
		v := calc(s[p+1 : e-1])

		s = s[:p-1] + fmt.Sprintf(" %d ", v) + s[e+1:]
		for strings.Index(s, "(") > -1 {
			s = fmt.Sprintf("%d", calc(s))
		}
	}
	s = strings.TrimSpace(s)
	v := strings.Split(s, " ")
	num := c.GetInt(v[0])
	s = ""
	for i := 1; i < len(v); i += 2 {
		if v[i] == "+" {
			num += c.GetInt(v[i+1])
		} else if v[i] == "*" {
			s += fmt.Sprintf("%d * ", num)
			num = c.GetInt(v[i+1])
		}
	}
	if num > 0 {
		s += fmt.Sprintf("%d", num)
	}
	v = strings.Split(s, " ")
	num = c.GetInt(v[0])
	for i := 1; i < len(v); i += 2 {
		if v[i] == "*" {
			num *= c.GetInt(v[i+1])
		}
	}
	return num
}

func main() {
	sum := 0

	for _, v := range c.ReadInputFile() {
		d := calc(v)
		//	c.Print(d)
		sum += d
	}
	c.Print(sum)

}
