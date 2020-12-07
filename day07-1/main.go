package main

import (
	c "adventofcode2020/common"
	"regexp"
	"strings"
)

var db map[string][]string

func search(target, start string) bool {
	for _, v := range db[start] {
		if v == target {
			return true
		}
		if search(target, v) {
			return true
		}
	}
	return false
}

func main() {
	db = map[string][]string{}
	sum := 0
	re := regexp.MustCompile(`\d\s([a-z]+\s[a-z]+)\sbag`)

	for _, v := range c.ReadInputFile() {
		g := strings.Split(v, " bags contain ")
		db[g[0]] = []string{}
		if g[1] == "no other bags." {
			continue
		}
		for _, v := range re.FindAllStringSubmatch(g[1], -1) {
			db[g[0]] = append(db[g[0]], v[1])
		}
	}
	for k := range db {
		if search("shiny gold", k) {
			sum++
		}
	}
	c.Print(sum)

}
