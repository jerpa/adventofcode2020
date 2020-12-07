package main

import (
	c "adventofcode2020/common"
	"regexp"
	"strings"
)

type bag struct {
	name string
	num  int
}

var db map[string][]bag

func search(start string) int {
	res := 1
	for _, v := range db[start] {
		res += v.num * search(v.name)
	}
	return res
}

func main() {
	db = map[string][]bag{}

	re := regexp.MustCompile(`(\d+)\s([a-z]+\s[a-z]+)\sbag`)

	for _, v := range c.ReadInputFile() {
		g := strings.Split(v, " bags contain ")
		db[g[0]] = []bag{}
		if g[1] == "no other bags." {
			continue
		}
		for _, v := range re.FindAllStringSubmatch(g[1], -1) {
			num := c.GetInt(v[1])
			db[g[0]] = append(db[g[0]], bag{name: v[2], num: num})
		}
	}

	c.Print(search("shiny gold") - 1)

}
