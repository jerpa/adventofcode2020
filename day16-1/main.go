package main

import (
	c "adventofcode2020/common"
	"strings"
)

type rule struct {
	min int
	max int
}

func main() {
	sum := 0
	rules := map[string][]rule{}

	mode := "rules"

	for _, v := range c.ReadInputFile() {
		if len(v) == 0 {
			continue
		}
		if v == "your ticket:" {
			mode = "ticket"
		} else if v == "nearby tickets:" {
			mode = "nearby"
		} else {
			if mode == "rules" {
				d := strings.Split(v, ": ")
				rules[d[0]] = []rule{}
				for _, r := range strings.Split(d[1], " or ") {
					minMax := strings.Split(r, "-")
					rules[d[0]] = append(rules[d[0]], rule{min: c.GetInt(minMax[0]), max: c.GetInt(minMax[1])})
				}
			} else if mode == "ticket" {

			} else if mode == "nearby" {
			mainLoop:
				for _, t := range c.GetInts(strings.Split(v, ",")) {
					for _, r := range rules {
						for _, p := range r {
							if p.min <= t && t <= p.max {
								continue mainLoop
							}
						}
					}
					sum += t
				}
			}
		}

	}
	c.Print(sum)

}
