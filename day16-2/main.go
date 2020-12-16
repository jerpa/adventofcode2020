package main

import (
	c "adventofcode2020/common"
	"strings"
)

type rule struct {
	min int
	max int
}

type classValidator struct {
	rules       *[]rule
	possiblePos []int
}

func (c classValidator) valid(v int) bool {
	for _, r := range *c.rules {
		if r.min <= v && v <= r.max {
			return true
		}
	}
	return false
}
func (c *classValidator) addRule(r rule) {
	if c.rules == nil {
		c.rules = &[]rule{}
	}
	*c.rules = append(*c.rules, r)
}
func (c *classValidator) removePos(pos int) bool {
	for i := len(c.possiblePos) - 1; i >= 0; i-- {
		if c.possiblePos[i] == pos {
			c.possiblePos = append(c.possiblePos[:i], c.possiblePos[i+1:]...)
			return true
		}
	}
	return false
}

func main() {
	sum := 0
	rules := map[string]*classValidator{}
	ticket := []int{}
	mode := "rules"
	rulePos := map[string]int{}
	validTickets := [][]int{}
mainLoop:
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
				rules[d[0]] = &classValidator{}
				for _, r := range strings.Split(d[1], " or ") {
					minMax := strings.Split(r, "-")
					rules[d[0]].addRule(rule{min: c.GetInt(minMax[0]), max: c.GetInt(minMax[1])})
				}
			} else if mode == "ticket" {
				ticket = c.GetInts(strings.Split(v, ","))
				validTickets = append(validTickets, c.GetInts(strings.Split(v, ",")))
			} else if mode == "nearby" {

				for _, t := range c.GetInts(strings.Split(v, ",")) {
					ok := false
					for _, r := range rules {
						if r.valid(t) {
							ok = true
						}
					}
					if !ok {
						continue mainLoop
					}
				}
				validTickets = append(validTickets, c.GetInts(strings.Split(v, ",")))
			}
		}
	}
	// Find possible positions
	for r, v := range rules {
	ticketLoop:
		for i := 0; i < len(validTickets[0]); i++ {
			for _, t := range validTickets {
				if !v.valid(t[i]) {
					continue ticketLoop
				}
			}
			rules[r].possiblePos = append(rules[r].possiblePos, i)
		}
	}
	// Remove duplicates
	for {
		ok := true
		for k := range rules {
			if len(rules[k].possiblePos) == 1 {
				for rk := range rules {
					if rk == k {
						continue
					}
					if rules[rk].removePos(rules[k].possiblePos[0]) {
						ok = false
					}
				}
			}
		}
		if ok {
			break
		}
	}
	sum = 1
	for k := range rules {
		if strings.HasPrefix(k, "departure") {
			sum *= ticket[rules[k].possiblePos[0]]
		}
	}
	c.Print(rulePos, sum, ticket)

}
